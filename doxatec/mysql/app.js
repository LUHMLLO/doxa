var mysql = require("mysql");
var conexion = mysql.createConnection({
  host: "142.93.207.120",
  database: "Smart_Nevera",
  user: "Nevera",
  password: "0144250809JDR@f",
});

conexion.connect(function (error) {
  if (error) {
    throw error;
  } else {
    console.log("Conexion Exitosa");
  }
});

// conexion.query(
//   "INSERT INTO `TempReg` (`DeviceName`, `TempSuperior`, `TempIntermedia`, `TempInferior`, `Usuario`) VALUES ('DevClient - Smart Fridge #01', 9, 6, 3, 'xd0asr7mbu8vq2v')"
// );

conexion.query("SELECT * from TempReg", function (error, results, fields) {
  if (error) {
    throw error;
  }
  results.forEach((result) => {
    console.log(result);
  });
});

conexion.end();
