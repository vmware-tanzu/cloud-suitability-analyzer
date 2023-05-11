ActivityBuilder builder = new ActivityBuilder
{
   Name = workflowName,
   Implementation = this.workflowImpl,
};

using (var fs = XmlWriter.Create(WorkflowFilename, settings))
{
   using (var xmlW = new XamlXmlWriter(fs, new XamlSchemaContext()))
   {
      using (var xw = System.Activities.XamlIntegration.ActivityXamlServices.CreateBuilderWriter(xmlW))
         {
            XamlServices.Save(xw, builder);
         }
      }
   }
}