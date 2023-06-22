using System;
using System.Collections;
using System.Collections.Generic;
using System.Text;
using System.Data;
using FirebirdSql.Data.FirebirdClient;
using System.Diagnostics;
using System.Windows.Forms;
using System.Threading;
using System.Runtime.Remoting;
using System.Runtime.Remoting.Channels;
using System.Runtime.Remoting.Channels.Ipc;
using System.Security.Permissions;
using System.Security.Cryptography;
using System.IO;
using Microsoft.Win32;

namespace ABCAddressBook
{
    public static class ABCUtilities
    {
        // Only Change these. #############
        public static bool SingleUserVersion = false;
        private static string ConnectionStringLocal = "User=$username;Password=$password;Database=$database;ServerType=1;Dialect=3;Charset=ISO8859_1;";
        private static string ConnectionStringServer = "User=$username;Password=$password;Database=$database;ServerType=0;Server=$server;Port=3050;Dialect=3;Charset=ISO8859_1;";        
        // ################################
                
        private static string _username;
        private static string _password;        
        private static FbRemoteEvent _remoteEvents;
        public static FbConnection DataConnection; // Persistant database connection.        
        private static MainForm mainForm;

        private static ABCRemoteObject _remoteObject = new ABCRemoteObject();
        private static IpcChannel _ipcChannel;
        public static bool UsingServer = false;
        public static bool Registered = false;
        
        private static string GetConnectionString()
        {
            string ConnectionString = "";
            if (SingleUserVersion) { SetRegistryValue("ConnectionType", "Local File"); }

            switch (GetRegistryValue("ConnectionType"))
            {
                case "Server":
                    ConnectionString = ConnectionStringServer;
                    UsingServer = true;
                    break;
                case "Local File":
                    ConnectionString = ConnectionStringLocal;
                    UsingServer = false;
                    break;
                default:
                    return "error";
            }

            if (SingleUserVersion)
            {
                ConnectionString = ConnectionString.Replace("$username", "SYSDBA");
                ConnectionString = ConnectionString.Replace("$password", "masterkey");
                ConnectionString = ConnectionString.Replace("$database", System.Environment.GetFolderPath(Environment.SpecialFolder.CommonApplicationData) + "\\" + Application.CompanyName + "\\" + Application.ProductName + "\\" + "ABCAddressBook.fdb");
                ConnectionString = ConnectionString.Replace("$server", "");
            }
            else
            {
                ConnectionString = ConnectionString.Replace("$username", _username);
                ConnectionString = ConnectionString.Replace("$password", _password);
                ConnectionString = ConnectionString.Replace("$database", GetRegistryValue("Database"));
                ConnectionString = ConnectionString.Replace("$server", GetRegistryValue("Server"));
            }

            return ConnectionString;
        }

        public static string EscapeString(string escapeString)
        {
            escapeString = escapeString.Replace("'", "''");
            return escapeString;
        }

        public static bool Connect()
        {            
            Boolean loginNotFinished = true;            
            Boolean notConnected = true;

            do
            {
                notConnected = true;
                do
                {
                    loginNotFinished = true;

                    if ((GetRegistryValueBool("AutoLogin")) || (SingleUserVersion))
                    {                        
                        _username = GetRegistryValue("UserName").ToUpper();                        
                        _password = GetRegistryValue("Password");

                        loginNotFinished = false;
                    }
                    else
                    {
                        LoginResult result = LoginForm.DoLogin();
                        switch (result.SelectedButton)
                        {
                            case SelectedLoginButton.Login:
                                _username = result.UserName.ToUpper();
                                _password = result.Password;
                                loginNotFinished = false;
                                break;

                            case SelectedLoginButton.Exit:
                                return false;

                            case SelectedLoginButton.Properties:
                                DatabasePropertiesForm i = new DatabasePropertiesForm();
                                i.ShowDialog();
                                break;
                        }
                    }
                } while (loginNotFinished);
                
                DataConnection = new FbConnection(ABCUtilities.GetConnectionString());
                try { DataConnection.Open(); }
                catch (FbException ex)
                {
                    ABCLog.LogWriteFbException(ex);
                    ABCLog.LogWriteText(ABCUtilities.GetConnectionString());
                    switch (ex.ErrorCode)
                    {
                        case -2147467259:
                            MessageBox.Show("Invalid username or password! - " + ex.Message, Application.ProductName);
                            if (SingleUserVersion) { return false; }

                            if (GetRegistryValueBool("AutoLogin")) 
                            {
                                SetRegistryValueBool("AutoLogin", false);                                
                            }
                            
                            break;

                        default:
                            MessageBox.Show("There was an error connecting to the database, please check the database settings and restart the program. MESSAGE: " + ex.Message + " Error Number: " + ex.ErrorCode);

                            DatabasePropertiesForm i = new DatabasePropertiesForm();
                            DialogResult result = i.ShowDialog();
                            if (result == DialogResult.Cancel) { return false; }                                                        
                            break;
                    }
                }
                
                if (DataConnection.State == System.Data.ConnectionState.Open) 
                {
                    // setup remote event handling.
                    if (UsingServer)
                    {
                        _remoteEvents = new FbRemoteEvent(DataConnection, new string[] { "MESSAGE_" + _username });
                        _remoteEvents.RemoteEventCounts += new FbRemoteEventEventHandler(RemoteEvent);
                        _remoteEvents.QueueEvents();
                    }
                   
                    // IPC Remoting
                    _ipcChannel = new IpcChannel("ABCAddressBook");
                    ChannelServices.RegisterChannel(_ipcChannel, false);                                        
                    RemotingServices.Marshal(_remoteObject, "ABCRemoteServer");
                    _remoteObject.CommandsEvent += new ABCRemoteObject.CommandsHandler(RemoteCommands);                    

                    // Set the default username to the current username.
                    if (!GetRegistryValueBool("AutoLogin"))
                    {
                        SetRegistryValue("UserName", _username);
                    }

                    // Setup state monitoring.
                    DataConnection.StateChange += DataConnection_StateChange;

                    notConnected = false;
                }
            } while (notConnected);
            
            return true;
        }

        public static string GetUserName()
        {
            return _username;
        }

        private static void RemoteEvent(object sender, FbRemoteEventEventArgs e)
        {            
            if (e.Name == "MESSAGE_" + _username)
            {
                GetMessages();
            }
        }

        public static void GetMessages()
        {            
            FbCommand fbcmd = ABCUtilities.DataConnection.CreateCommand();
            fbcmd.CommandText = "SELECT MESSAGE_ID FROM MESSAGES WHERE MESSAGE_TO = @USERNAME AND MESSAGE_DISPLAYED = 0;";
            fbcmd.Parameters.Add("@USERNAME", _username);
            FbDataReader reader = fbcmd.ExecuteReader();
            while (reader.Read())
            {
                if (mainForm != null)
                {
                    if (mainForm.DisplayMessage != null)
                    { mainForm.Invoke(mainForm.DisplayMessage, reader.GetInt16(reader.GetOrdinal("MESSAGE_ID"))); }
                }
            }
        }

        public static void GetUnreadMessages()
        {
            FbCommand fbcmd = ABCUtilities.DataConnection.CreateCommand();
            fbcmd.CommandText = "SELECT MESSAGE_ID FROM MESSAGES WHERE MESSAGE_TO = @USERNAME AND MESSAGE_DISPLAYED = 1 AND MESSAGE_READ = 0;";
            fbcmd.Parameters.Add("@USERNAME", _username);
            FbDataReader reader = fbcmd.ExecuteReader();
            while (reader.Read())
            {
                if (mainForm != null)
                {
                    if (mainForm.DisplayMessage != null)
                    { mainForm.Invoke(mainForm.DisplayMessage, reader.GetInt16(reader.GetOrdinal("MESSAGE_ID"))); }
                }
            }
        }

        public static DateTime GetServerTime()
        {
            FbCommand fbcmd = ABCUtilities.DataConnection.CreateCommand();
            fbcmd.CommandType = CommandType.StoredProcedure;              
            fbcmd.CommandText = "GET_TIMESTAMP";
            FbParameter result = fbcmd.Parameters.Add("@CURRENTTIMESTAMP", FbDbType.TimeStamp, 1);
            result.Direction = ParameterDirection.Output;
            fbcmd.ExecuteNonQuery();            

            return (DateTime)result.Value;
        }

        public static string GetScreenName(string username)
        {
            FbCommand fbcmd = ABCUtilities.DataConnection.CreateCommand();
            fbcmd.CommandType = CommandType.StoredProcedure;
            fbcmd.CommandText = "GET_SCREENNAME";
            
            fbcmd.Parameters.Add("@USERNAME", username);
            
            FbParameter result = fbcmd.Parameters.Add("@SCREENNAME", FbDbType.VarChar, 30);
            result.Direction = ParameterDirection.Output;
            fbcmd.ExecuteNonQuery();            

            return (string)result.Value;            
        }

        public static void SetMessageDisplayed(int messageID)
        {
            FbCommand fbcmd = ABCUtilities.DataConnection.CreateCommand();
            fbcmd.CommandType = CommandType.StoredProcedure;
            fbcmd.CommandText = "SET_MESSAGE_DISPLAYED";
            fbcmd.Parameters.Add("@MESSAGE_ID", messageID);            
            fbcmd.ExecuteNonQuery();
            return;
        }

        public static void SetMessageRead(int messageID)
        {
            FbCommand fbcmd = ABCUtilities.DataConnection.CreateCommand();
            fbcmd.CommandType = CommandType.StoredProcedure;
            fbcmd.CommandText = "SET_MESSAGE_READ";
            fbcmd.Parameters.Add("@MESSAGE_ID", messageID);
            fbcmd.ExecuteNonQuery();
            return;
        }

        private static void RemoteCommands(string[] args)
        {
            mainForm.BeginInvoke(mainForm.RemoteCommandsEvent, new object[] { args }); 
        }
        
        [STAThread]
        static void Main(string[] args)
        {            
            string[] a = { "Open" };
            Process[] p = Process.GetProcessesByName(Process.GetCurrentProcess().ProcessName);
            if (p.Length > 1)
            {                
                ABCRemoteObject remObject = (ABCRemoteObject)Activator.GetObject(typeof(ABCRemoteObject), "ipc://ABCAddressBook/ABCRemoteServer");
                if (remObject == null) 
                {
                    MessageBox.Show("Could not get remote IPC object.", Application.ProductName);
                    ABCLog.LogWriteText("Could not get remote IPC object.");
                }
                else 
                {
                    try
                    { remObject.Commands(a); }
                    catch
                    {
                        MessageBox.Show("Could not send commands to remote IPC object.", Application.ProductName);
                        ABCLog.LogWriteText("Could not send commands to remote IPC object.");
                    }
                }
            }
            else
            {
                for (int i = 0; i < args.Length; i++)
                {
                    switch (args[i])
                    {
                        case "hideonstart":                            
                            SetRegistryValueBool("HideOnStart", true);
                            break;
                        default: break;
                    }
                }

                Application.EnableVisualStyles();
                Application.SetCompatibleTextRenderingDefault(false);
                Thread.CurrentThread.Name = "Main Thread";                

                if (CheckRegistration())
                {
                    if (Connect())
                    {
                        Application.Run(mainForm = new MainForm());
                    }
                }
            }
        }

        private static bool CheckRegistration()
        {
            bool regCheck = false;
            bool regPass = false;
            TimeSpan trialDuration = new TimeSpan(30, 0, 0, 0);
            TimeSpan currentTime = new TimeSpan(DateTime.Now.Ticks);
            
            while (!regCheck)
            {
                if (GetRegistrationName() != "")
                {
                    if (GetRegistrationKey() == Encode(GetRegistrationName()))
                    {
                        GetTrialTime();
                        Registered = true;
                        regCheck = true;
                        regPass = true;
                    }
                    else
                    {
                        MessageBox.Show("Your registration settings are invalid, please revise.", "ABC Address Book");
                        RegisterForm rf = new RegisterForm();
                        DialogResult dr = rf.ShowDialog();
                        if (dr == DialogResult.Cancel)
                        { 
                            regPass = false;
                            regCheck = true;
                        }
                    }
                }
                else
                {
                    TimeSpan trialTime = new TimeSpan(GetTrialTime());
                    if (trialTime.Ticks == 0) 
                    {
                        ABCLog.LogWriteText("Registration details have been tampered with!");
                        MessageBox.Show("Naughty Naughty Naughty", "ABC Address Book");
                        regCheck = true;
                        regPass = false;
                        break;
                    }

                    TimeSpan expireTime = new TimeSpan(trialTime.Ticks);
                    expireTime = expireTime.Add(trialDuration);                    

                    if (currentTime > expireTime)
                    {
                        MessageBox.Show("Your 30 day trial peroid is over, if you wish to continue using this product, please purchase a registration key.", "ABC Address Book");
                        RegisterForm rf = new RegisterForm();
                        DialogResult dr = rf.ShowDialog();
                        if (dr == DialogResult.Cancel)
                        {
                            regPass = false;
                            regCheck = true;
                        }                        
                    }
                    else
                    {
                        TimeSpan displacment = new TimeSpan(expireTime.Ticks);
                        TimeSpan zeroTime = new TimeSpan(0);
                        displacment = displacment.Subtract(currentTime);
                        if (displacment > zeroTime && displacment <= trialDuration)
                        {
                            MessageBox.Show("You have " + displacment.Days.ToString() + " days left in your trial period.", "ABC Address Book");
                            regCheck = true;
                            regPass = true;
                        }
                        else
                        {
                            ABCLog.LogWriteText("Registration details have been tampered with!  Displacement = " + displacment.Days.ToString());
                            MessageBox.Show("Naughty Naughty Naughty", "ABC Address Book");
                            regCheck = true;
                            regPass = false;
                        }
                    }
                }
            }
            return regPass;
        }

        public static void Quit()
        {
            ChannelServices.UnregisterChannel(_ipcChannel);
            Application.Exit();
        }
      
        public static string Encode(string name)
        {
            name = name.PadRight(16, ' ');
            Random random = new Random();            
            string text = "";
            
            for (int i = 0; i < name.Length; i++)
            {
                if (i == 0) { text += ToBase36(name[i] ^ 20); }
                else { text += ToBase36(name[i] ^ name[i - 1]); }
            }
            
            return text.Remove(0, text.Length - 16);
        }

        public static string ToBase36(int num)
        {
            string base36num = "";
            int remainder;
            int division;

            do
            {
                remainder = num % 36;
                division = num / 36;
                base36num = ToBase36Char(remainder).ToString() + base36num;                
                num = division;
            } while (division != 0);
            
            return base36num;
        }

        public static char ToBase36Char(int num)
        {
            string alpha = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ";
            return alpha[num];
        }

        public static int Base36CharToDec(char numeral)
        {            
            int num = Convert.ToInt16(numeral);
            if (num >= 48 && num <= 57) { return num - 48; }
            if (num >= 65 && num <= 90) { return num - 55; }
            return 0;
        }

        public static string GetDataConfig(string configName)
        {
            FbCommand fbcmd = ABCUtilities.DataConnection.CreateCommand();
            fbcmd.CommandType = CommandType.StoredProcedure;
            fbcmd.CommandText = "READ_CONFIG";
            fbcmd.Parameters.Add("@CNAME", configName);
            FbParameter result = fbcmd.Parameters.Add("@CVALUE", FbDbType.VarChar, 255);
            result.Direction = ParameterDirection.Output;
            fbcmd.ExecuteNonQuery();

            return result.Value.ToString();
        }

        public static void SaveDataConfig(string configName, string configValue)
        {
            FbCommand fbcmd = ABCUtilities.DataConnection.CreateCommand();
            fbcmd.CommandType = CommandType.StoredProcedure;
            fbcmd.CommandText = "SAVE_CONFIG";
            fbcmd.Parameters.Add("@CNAME", configName);
            fbcmd.Parameters.Add("@CVALUE", configValue);
            fbcmd.ExecuteNonQuery();
        }

        public static string GetRegistrationName()
        { return GetRegistryValue("RegistrationName"); }

        public static string GetRegistrationKey()
        { return GetRegistryValue("RegistrationKey"); }

        public static void SetRegistrationName(string name)
        { SetRegistryValue("RegistrationName", name); }

        public static void SetRegistrationKey(string rKey)
        { SetRegistryValue("RegistrationKey", rKey); }

        public static long GetTrialTime()
        {
            RegistryKey key = GetRegistryKey();
            
            string first = (string)key.GetValue("ABC", null);
            long tickLong = Convert.ToInt64(key.GetValue("RegTicks", 0));

            if (tickLong == 0 && first == "first")            
            {                
                key.SetValue("ABC", "34556");
                long ticks = DateTime.Now.Ticks;
                key.SetValue("RegTicks", ticks, RegistryValueKind.QWord);
                tickLong = ticks;
            }
            else if (first == "first") { key.SetValue("ABC", "34556"); }
            return tickLong;
        }

        public static string GetRegistryValue(string name, string sdefault)
        {
            RegistryKey key = GetRegistryKey();
            return (string)(key.GetValue(name, sdefault));
        }

        public static string GetRegistryValue(string name)
        {
            RegistryKey key = GetRegistryKey();
            return (string)(key.GetValue(name, ""));
        }

        public static void SetRegistryValue(string name, string value)
        {
            RegistryKey key = GetRegistryKey();
            key.SetValue(name, value);
        }

        public static bool GetRegistryValueBool(string name)
        {
            RegistryKey key = GetRegistryKey();
            int i = (int)(key.GetValue(name, 0));
            if (i == 0) { return false; } else { return true; }            
        }

        public static void SetRegistryValueBool(string name, bool value)
        {
            RegistryKey key = GetRegistryKey();            
            if (value)
            { key.SetValue(name, 1, RegistryValueKind.DWord); }
            else { key.SetValue(name, 0, RegistryValueKind.DWord); }
        }

        public static RegistryKey GetRegistryKey()
        {
            RegistryKey key = Registry.LocalMachine.OpenSubKey("Software\\SFS\\ABCAddressBook", true);
            // If the return value is null, the key doesn't exist
            if (key == null)
            {
                // The key doesn't exist; create it / open it
                key = Registry.LocalMachine.CreateSubKey("Software\\" + Application.CompanyName + "\\" + Application.ProductName, RegistryKeyPermissionCheck.ReadWriteSubTree);
            }

            return key;
        }

        public static bool IsUserOnline(string user)
        {
            FbDatabaseInfo dinfo = new FbDatabaseInfo(ABCUtilities.DataConnection);
            ArrayList ulist = dinfo.ActiveUsers;
            if (ulist.Contains(user)) { return true; } else { return false; }
        }

        public static void DataConnection_StateChange(object sender, System.Data.StateChangeEventArgs e)
        {
            // Do nothing, does not respond soon enough.
            switch (e.CurrentState)
            {
                case ConnectionState.Broken:
                    MessageBox.Show("Connection to server is broken.", Application.ProductName);
                    ABCUtilities.Quit();
                    break;
                case ConnectionState.Closed:
                    MessageBox.Show("Connection to server has closed.", Application.ProductName);
                    ABCUtilities.Quit();
                    break;
                default:
                    break;
            }
        }    
    }
}
