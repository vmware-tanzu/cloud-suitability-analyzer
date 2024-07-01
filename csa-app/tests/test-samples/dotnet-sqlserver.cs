using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;
using System.Data.OleDb;
using System.Data.Common;
using System.Data.SqlClient;
using Microsoft.SqlServer;
using Microsoft.SqlServer.Management.Smo;

namespace ExcelToSQL
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        [System.Runtime.InteropServices.DllImport("user32.dll")]
        private static extern IntPtr SendMessage(IntPtr hWnd, int msg, IntPtr wp, IntPtr lp);

        private void button3_Click(object sender, EventArgs e)
        {
            // prepare the open file dialog
            OpenFileDialog ofd = new OpenFileDialog();
            ofd.CheckFileExists = true;
            ofd.CheckPathExists = true;
            ofd.AutoUpgradeEnabled = true;
            ofd.Filter = "Excel Spreadsheet (*.xls)|*.xls";
            ofd.Multiselect = false;
            ofd.Title = "Select Excel Spreadsheet...";

            // launch the open file dialog
            if (ofd.ShowDialog() != DialogResult.Cancel)
            {
                this.textBox1.Text = ofd.FileName;
            }
        }

        private void Form1_Load(object sender, EventArgs e)
        {

        }

        private void button2_Click(object sender, EventArgs e)
        {
            // exit the application
            Application.Exit();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            // validate the user input
            if (this.textBox1.Text.Length == 0 ||
                this.textBox2.Text.Length == 0 ||
                this.textBox3.Text.Length == 0)
            {
                MessageBox.Show("Please ensure that you have completed all fields before clicking Import. Please check your input and try again.", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                return;
            }

            Application.UseWaitCursor = true;
            SendMessage(this.Handle, 0x20, this.Handle, (IntPtr)1);
            Application.DoEvents();

            // parse the connection string
            string sqlConnection = this.textBox2.Text;
            string database = sqlConnection.Split(new string[] { "database=" }, StringSplitOptions.None)[1].Split(Convert.ToChar(";"))[0];
            string server = sqlConnection.Split(new string[] { "server=" }, StringSplitOptions.None)[1].Split(Convert.ToChar(";"))[0];
            string uid = sqlConnection.Split(new string[] { "uid=" }, StringSplitOptions.None)[1].Split(Convert.ToChar(";"))[0];
            string pwd = sqlConnection.Split(new string[] { "pwd=" }, StringSplitOptions.None)[1].Split(Convert.ToChar(";"))[0];

            // connect to excel workbook
            string excelConnection = string.Format("Provider=Microsoft.Jet.OLEDB.4.0;Data Source={0};Extended Properties=Excel 8.0;", this.textBox1.Text);

            // execute procedure
            try
            {
                // open excel workbook
                using (OleDbConnection connection = new OleDbConnection(excelConnection))
                {
                    connection.Open();

                    DataTable tableSchema = connection.GetOleDbSchemaTable(OleDbSchemaGuid.Tables, new Object[] { null, null, null, "TABLE" });
                    string table = tableSchema.Rows[0]["Table_Name"].ToString();

                    OleDbCommand command = new OleDbCommand(string.Format("SELECT * FROM [{0}]", table), connection);

                    using (DbDataReader dr = command.ExecuteReader())
                    {
                        // create new table
                        Microsoft.SqlServer.Management.Common.ServerConnection smoConn = new Microsoft.SqlServer.Management.Common.ServerConnection();
                        smoConn.ServerInstance = server;
                        smoConn.LoginSecure = false;
                        smoConn.Login = uid;
                        smoConn.Password = pwd;
                        Server dbserver = new Server(smoConn);
                        Database db = dbserver.Databases[database];
                        DataTable schema = dr.GetSchemaTable();
                        Table newTable = new Table(db, this.textBox3.Text);
                        for (int i = 0; i < schema.Rows.Count; i++)
                        {
                            Column column = new Column(newTable, schema.Rows[i][0].ToString());
                            column.DataType = DataType.VarChar(5000);
                            column.Nullable = true;
                            newTable.Columns.Add(column);
                        }
                        newTable.Create();

                        // copy data
                        using (SqlBulkCopy bulkCopy = new SqlBulkCopy(sqlConnection))
                        {
                            bulkCopy.DestinationTableName = this.textBox3.Text;
                            bulkCopy.WriteToServer(dr);
                        }
                    }
                }

                MessageBox.Show("Excel Spreadsheet data was imported successfully!", "Information", MessageBoxButtons.OK, MessageBoxIcon.Information);
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message, "Error", MessageBoxButtons.OK, MessageBoxIcon.Error);
            }

            Application.UseWaitCursor = false;
            SendMessage(this.Handle, 0x20, this.Handle, (IntPtr)1);
            Application.DoEvents();
        }

        private void button4_Click(object sender, EventArgs e)
        {
            Form2 dialog = new Form2();
            if (dialog.ShowDialog() == DialogResult.OK)
            {
                this.textBox2.Text = dialog.ConnectionString;
            }
        }
    }
}
