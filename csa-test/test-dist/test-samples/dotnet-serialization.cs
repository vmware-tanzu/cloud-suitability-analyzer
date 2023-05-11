using System;
using System.Runtime.Serialization;
using System.IO;
using System.Runtime.Serialization.Formatters.Binary;
using System.Xml.Serialization;

public sealed class Program {
    public static void Main() {

        // create a file to save the data to
        FileStream fs = new FileStream("SerializedString.Date", FileMode.Open);

        // create an XmlSerializer object to perform the serialization
        XmlSerializer xs = new XmlSerializer(typeof(DateTime));

        // use the XmlSerializer object  to serialize the data to a file
        DateTime previousTime = (DateTime)xs.Deserialize(fs);

        // close the file
        fs.Close();

        // display the deserialized time
        Console.WriteLine("Day: " + previousTime.DayOfWeek + 
                " ,Time: " + previousTime.TimeOfDay.ToString());
    }
}
