using System.Web.Script.Serialization;

namespace console_client
{
    public static class JsonHelper
    {
        public static string ToJson(this object obj)
        {
            var serializer = new JavaScriptSerializer();
            return serializer.Serialize(obj);
        }

        public static string ToJson(this object obj, int recursionDepth)
        {
            var serializer = new JavaScriptSerializer();
            serializer.RecursionLimit = recursionDepth;
            return serializer.Serialize(obj);
        }
    }
    
    public class ProjectData
    {
        public string Id { get; set; }
        public string Name { get; set; }
        public string Desc { get; set; }
        public string Model { get; set; }

        public ProjectData(string id, string name, string desc, string model)
        {
            Id = id;
            Name = name;
            Desc = desc;
            Model = model;
        }
    }

    public class ComponentData
    {
        public string Id { get; set; } // Object id
        public string Name { get; set; } // Object name
        public string Profile { get; set; } // Object profile
        public string Material { get; set; } // Object material
        public string Class { get; set; } // Object class
        public int Co2 { get; set; } // Object co2/m2
        public string Phase { get; set; } // Production phases

        public ComponentData(string id, string name, string profile, string material, string @class, int co2, string phase)
        {
            Id = id;
            Name = name;
            Profile = profile;
            Material = material;
            Class = @class;
            Co2 = co2;
            Phase = phase;
        }
    }
}