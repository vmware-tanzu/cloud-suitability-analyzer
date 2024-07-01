using System;
using System.Data;
using System.Data.SqlClient;
using Npgsql;

namespace HealthMonitorSystem
{
	
	
	public class BaseDA
	{
		private static string connectionString = "Server=db.cecs.bdx.go;" + "Port=5432;" + 
								"Database=gulab;" + "User ID=jamun;" + "Password=yu!m11Ab;";
		private static NpgsqlConnection dbcon = new Npgsql.NpgsqlConnection(connectionString);
		
				
		public static int ExecuteNonQuery(string strSql)
		{
			try
			{
				dbcon.Open();
				NpgsqlCommand cmd = new NpgsqlCommand(strSql, dbcon);
				int retVal = cmd.ExecuteNonQuery();
				return retVal;			
			}
			catch
			{
				throw;
			}
			finally
			{
				dbcon.Close();
			}
		}
		
		public static int ExecuteScalar(string strSql)
		{
			try
			{
				dbcon.Open();
				NpgsqlCommand cmd = new NpgsqlCommand(strSql, dbcon);
				Object retVal = cmd.ExecuteScalar();
				dbcon.Close();
				return Convert.ToInt32(retVal);			
			}
			catch
			{
				throw;
			}
			finally
			{
				dbcon.Close();
			}
		}
		
		public static DataSet ExecuteDataSet(string strSql)
		{			
			try
			{
				NpgsqlDataAdapter da = new NpgsqlDataAdapter(strSql, connectionString);
				DataSet ds = new DataSet();
				da.Fill(ds); 				
				return ds;
			}
			catch
			{	
				throw;
			}
			finally
			{
				dbcon.Close();
			}
			
		}
		
		public static string Escape(string strStmt)
		{
			strStmt = strStmt.Trim();
			strStmt = strStmt.Replace("'", "''");
			return strStmt;
		}
		
	}
}
