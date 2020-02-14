using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Threading.Tasks;

namespace Web.Helpers
{
    public class HttpClientHelper
    {
        private static readonly string BaseUrl = "https://localhost:5001/api/v1/project";
        public static HttpClient GetHttpClient()
        {
            var client = new HttpClient();
            client.BaseAddress = new Uri(BaseUrl);
            return client;
        }
    }
}
