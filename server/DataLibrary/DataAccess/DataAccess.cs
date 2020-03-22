using System;
using System.Collections.Generic;
using Dapper;
using System.Data;
using System.Linq;
using Npgsql;

namespace DataLibrary.DataAccess
{
    public static class DataAccess
    {
        private static String connectionString;
        private static NpgsqlConnection GetConnection()
        {
            return new NpgsqlConnection(GetConnectionString());
        }
        /// <summary>
        /// Set the connection string for database connection
        /// </summary>
        /// <param name="conn">Connection string from config</param>
        public static void SetConnectionString(String conn)
        {
            connectionString = conn;
        }
        private static string GetConnectionString()
        {
            return connectionString;
        }

        /// <summary>
        /// LoadData will query data from database, with specifed sql query.
        /// </summary>
        /// <typeparam name="T">Type of reslut to cast</typeparam>
        /// <param name="sql">SQL query</param>
        /// <returns>return list<T> if list is null or empty no data got returned</returns>
        public static List<T> LoadData<T>(string sql)
        {
            using (IDbConnection cnn = GetConnection())
            {
                return cnn.Query<T>(sql).ToList();
            }
        }
        /// <summary>
        /// LoadData will query data from database, with specifed sql query.
        /// </summary>
        /// <typeparam name="T">Type of reslut to cast</typeparam>
        /// <param name="sql">SQL query</param>
        /// <param name="data">Paramters in query</param>
        /// <returns>return list<T> if list is null or empty no data got returned</returns>
        public static List<T> LoadData<T>(string sql, string data)
        {
            using (IDbConnection cnn = GetConnection())
            {
                return cnn.Query<T>(sql, new { data }).ToList();
            }
        }

        /// <summary>
        /// SaveData will be called when you want to save data to database
        /// </summary>
        /// <typeparam name="T">Type of data</typeparam>
        /// <param name="sql">SQL Query</param>
        /// <param name="data">Data you want to store</param>
        /// <returns>Returns total effected rows</returns>
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

        public static void SeedDatabase()
        {

        }
    }
}
