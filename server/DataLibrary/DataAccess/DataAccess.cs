﻿using System;
using System.Collections.Generic;
using System.Text;
using Dapper;
using System.Data;
using Microsoft.Extensions.Configuration;
using System.IO;
using System.Linq;
using Npgsql;

namespace DataLibrary.DataAccess
{
    public static class DataAccess
    {
        private static NpgsqlConnection GetConnection()
        {
            return new NpgsqlConnection(GetConnectionString());
        }
        public static string GetConnectionString()
        {
            IConfigurationRoot configuration = new ConfigurationBuilder()
                .SetBasePath(Directory.GetCurrentDirectory())
                .AddJsonFile(@Directory.GetCurrentDirectory() + "/appsettings.json")
                .Build();
            return configuration.GetConnectionString("DefaultConnection");
        }

        public static List<T> LoadData<T>(string sql)
        {
            using (IDbConnection cnn = GetConnection())
            {
                return cnn.Query<T>(sql).ToList();
            }
        }

        public static int SaveData<T>(string sql, T data)
        {
            using (IDbConnection cnn = GetConnection())
            {
                return cnn.Execute(sql, data);
            }
        }

        public static T QueryData<T>(string sql, T data)
        {
            using (IDbConnection cnn = GetConnection())
            {
                return cnn.QuerySingle<T>(sql, new { data });
            }
        }
    }
}