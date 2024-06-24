// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

using NUnit.Framework;
using Dotnet.Storm.Example;
using Dotnet.Storm.Adapter.Test;
using Dotnet.Storm.Adapter.Components;
using System.Collections.Generic;

namespace Dotnet.Storm.Test.Example
{
    public static class SmokeTests
    {
        [Test] public static void TestExample()
        {
            // Initialize configuration
            Dictionary<string, object> config = GetConfig();

            // Initialize context
            StormContext sc = GetContext();

            // Create, and run a spout
            EmitSentence es = UnitTest.CreateComponent<EmitSentence>(sc, config);
            es.Next();
            List<TestOutput> res = es.GetOutput();

            // Verify results and metadata
            Assert.True(res.Count > 0);
            Assert.True(res[0].Stream == "default");
            Assert.True(res[0].ComponentId == "componentid1");

            // Create, and run 1st Bolt
            sc.ComponentId = "componentid2";
            SplitSentence ss = UnitTest.CreateComponent<SplitSentence>(sc, config);
            foreach (var output in res)
            {
                ss.Execute(new StormTuple(((SpoutOutput)output).Id, "EmitSentence", "TaskId", output.Stream, output.Tuple));
            }
            res = ss.GetOutput();

            // Verify results and metadata
            Assert.True(res.Count > 0);
            Assert.True(res[0].Stream == "default");

            // Create, and run 2nd Bolt
            sc.ComponentId = "componentid3";
            CountWords cw = UnitTest.CreateComponent<CountWords>(sc, config);
            foreach (var output in res)
            {
                cw.Execute(new StormTuple("id", "SplitSentence", "TaskId", output.Stream, output.Tuple));
            }
        }

        private static Dictionary<string, object> GetConfig() => new Dictionary<string, object>
        {
            ["configkey1"] = "configvalue1",
            ["configkey2"] = "configvalue2"
        };

        private static StormContext GetContext() => new StormContext
        {
            ComponentId = "componentid1",
            StreamToOputputFields = new Dictionary<string, List<string>>() { { "default", new List<string>() { "title" } } }
        };
    }
}