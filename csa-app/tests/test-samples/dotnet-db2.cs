using IBM.Data.DB2;
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace Proyecto1_TBD2.Cheks {
    public partial class MostrarCheck:Form {
        TreeView arbol;
        public MostrarCheck(TreeView _arbol) {
            InitializeComponent();
            arbol = _arbol;
        }

        private void MostrarCheck_Load(object sender, EventArgs e) {
            PantallaPrincipal pn = new PantallaPrincipal();
            DB2Connection connection = pn.obtenerConexion(arbol.SelectedNode.Parent.Parent.Text);
            try {
                connection.Open();
                DB2Command cmd = new DB2Command(@"SELECT TEXT FROM SYSIBM.SYSCHECKS WHERE TBCREATOR='DB2ADMIN' AND NAME ='" + arbol.SelectedNode.Text + "';", connection);
                DB2DataReader buffer = cmd.ExecuteReader();
                while (buffer.Read()) {
                    var function = buffer ["TEXT"].ToString();
                    richTextBox1.Text = function;
                    break;
                }
                buffer.Close();
            } catch (DB2Exception ex) {
                MessageBox.Show("Error al mostrar Check\n" + ex.Message);
            }
            connection.Close();
        }
    }
}
