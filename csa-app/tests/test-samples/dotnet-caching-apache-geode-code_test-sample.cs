/*
* Licensed to the Apache Software Foundation (ASF) under one or more
* contributor license agreements.  See the NOTICE file distributed with
* this work for additional information regarding copyright ownership.
* The ASF licenses this file to You under the Apache License, Version 2.0
* (the "License"); you may not use this file except in compliance with
* the License.  You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

using System;
using Apache.Geode.Client;

namespace Apache.Geode.Examples.PutGetRemove
{
  class Program
  {
    static void Main(string[] args)
    {
      var cache = new CacheFactory()
          .Set("log-level", "none")
          .Create();

      cache.GetPoolManager()
          .CreateFactory()
          .AddLocator("localhost", 10334)
          .Create("pool");

      var regionFactory = cache.CreateRegionFactory(RegionShortcut.PROXY)
          .SetPoolName("pool");
      var region = regionFactory.Create<string, string>("example_userinfo");

      Console.WriteLine("Storing id and username in the region");

      const string rtimmonsKey = "rtimmons";
      const string rtimmonsValue = "Robert Timmons";
      const string scharlesKey = "scharles";
      const string scharlesValue = "Sylvia Charles";

      region.Put(rtimmonsKey, rtimmonsValue);
      region.Put(scharlesKey, scharlesValue);

      Console.WriteLine("Getting the user info from the region");
      var user1 = region.Get(rtimmonsKey, null);
      var user2 = region.Get(scharlesKey, null);

      Console.WriteLine(rtimmonsKey + " = " + user1);
      Console.WriteLine(scharlesKey + " = " + user2);

      Console.WriteLine("Removing " + rtimmonsKey + " info from the region");

      if (region.Remove(rtimmonsKey))
      {
        Console.WriteLine("Info for " + rtimmonsKey + " has been deleted");
      }
      else
      {
        Console.WriteLine("Info for " + rtimmonsKey + " has not been deleted");
      }

      cache.Close();
    }
  }
}