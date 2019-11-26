using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Newtonsoft.Json;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Text;

namespace console_client
{
    internal static class Program
    {
        public static void Main(string[] args)
        {
            var pd = new ProjectData("123", "Volvo ab", "This is an simple desc", "model");
            CreateProject(pd);
            GetProject(pd.Id);

            var csd = new List<ComponentData>{
                new ComponentData("1", "dunno", "4500*2431", "C#", "452", 3123, "Start"),
                new ComponentData("2", "dunno", "4500*2431", "C#2", "452", 3123, "Start")
            };
            
            AddComponentsToProject(csd, pd.Id);
            
            GetComponentsFromProject(pd.Id);
        }

        private static async Task<HttpResponseMessage> SendRequest(HttpMethod method, string endpoint, dynamic content = null)
        {
            HttpResponseMessage response = null;
            using (var client = new HttpClient())
            {
                using (var request = new HttpRequestMessage(method, endpoint))
                {
                    request.Headers.Accept.Add(new MediaTypeWithQualityHeaderValue("application/json"));

                    if (content != null)
                    {
                        string c;
                        if (content is string)
                            c = content;
                        else
                            c = JsonConvert.SerializeObject(content);
                        request.Content = new StringContent(c, Encoding.UTF8, "application/json");
                    }

                    response = await client.SendAsync(request).ConfigureAwait(false);
                }
            }
            
            return response;
        }

        private static void CreateProject(ProjectData data)
        {
            data.ToJson();
            try
            {
                var message = SendRequest(HttpMethod.Post, "http://localhost:5050/projects", data).Result;
                Console.WriteLine(message.Content.ReadAsStringAsync().Result);
                
                // Do something depending on status code and res data
                // Status codes that can be returned
                // 409 is returned when dup data for project is found: message will be "field: data, already exists" etc "id: 123, already exists"
                // 502 is also returned if error occur when inserting to db, but should be catch by above error
                // 201 is returned when it succeeded to create project with response of project object
            }
            catch (Exception e)
            {
                Console.WriteLine(e);
                throw;
            }
        }
        
        private static void GetProject(string projectId)
        {
            try
            {
                var message = SendRequest(HttpMethod.Get, "http://localhost:5050/projects/" + projectId).Result;
                Console.WriteLine(message.Content.ReadAsStringAsync().Result);
                
                // Do something depending on status code and res data
                // 404 will be returned if projectId passed will not be found in db message will be: "project not found"
                // 200 will be returned if project found with response object of project
            }
            catch (Exception e)
            {
                Console.WriteLine(e);
                throw;
            }
        }

        private static void AddComponentsToProject(List<ComponentData> data, string projectId)
        {
            try
            {
                var msg = SendRequest(HttpMethod.Post, "http://localhost:5051/components/" + projectId, data).Result;
                Console.WriteLine(msg.Content.ReadAsStringAsync().Result);
                
                // Do something depending on status code and res data
                // 404 will be returned if project id passed not found with message: "project not found"
                // 500 will be returned if it fails to add component to db. #Needs to be improved
                // 201 will be returned if it succeeded to add all of the components to db with message: "components added"
            }
            catch (Exception e)
            {
                Console.WriteLine(e);
                throw;
            }
        }
        
        private static void GetComponentsFromProject(string projectId)
        {
            try
            {
                var msg = SendRequest(HttpMethod.Get, "http://localhost:5051/components/" + projectId).Result;
                Console.WriteLine(msg.Content.ReadAsStringAsync().Result);
                // Do something depending on status code and res data
                // 404 will be returned if projectId passed will not be found in db message will be: "project not found"
                // 404 will also be returned if components for project does not exist message will be: "components not found"
                // 500 will be returned if error occur while grabbing data
                // 200 will be returned if project found with response object of an array with components
                
            }
            catch (Exception e)
            {
                Console.WriteLine(e);
                throw;
            }
        }
    }
}