using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Data.SqlClient;

namespace The_Gaming_Store
{
    public class db_con
    {
        public static SqlConnection getCon(){
        return new SqlConnection(System.Configuration.ConfigurationManager.ConnectionStrings["DB_TheGamingStoreConnectionString"].ConnectionString);
    }
    }
}