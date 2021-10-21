# RCP🚀
_Pagina web creada con Goolang y JavaScript para backend y fronted respectivamente, para la autenticacion de usuarios de la Red de Programacion Competitiva_

## Condiciones Tecnicas 🛠️

- Lenguajes de Programacion usados 💱 : Goolang,JavaScript
- Systemas Operativos usados 💻 : Windows 10 
- Entorno de desarrollo integrado utilizado 👨🏻‍💻 : Visual Studio Code, Visual Studio
- Instalacion 🔧 : Haga clic en Código, luego en Descargar ZIP

## Autores ✒️

* [Carolina Pasuy](https://github.com/CPASUY) - *Codificacion*  
* [Carlos Pantoja](https://github.com/CarlosJPantoja) - *Codificacion* 

## Expresiones de Gratitud 🎁

* Comenta a otros sobre este proyecto 📢
* Invita una cerveza 🍺 o un café ☕ a alguien del equipo. 
* Da las gracias públicamente 🤓.

## Como Usarlo💻
1. Descargar XAPP: Descargar y activar el MYSQL para utilizar la base de datos
https://www.apachefriends.org/es/download.html
2. Ingregar a la pagina del local host donde se encuentra las bases de datos lcoales
http://localhost/phpmyadmin/
4. Crear la base de datos rcp:
CREATE DATABASE rcp;
5. Crear la tabla usuario:
CREATE TABLE `usuarios` (
  `Username` varchar(30) NOT NULL,
  `Firstname` varchar(50) NOT NULL,
  `Lastname` varchar(50) NOT NULL,
  `Email` varchar(500) NOT NULL,
  `Password` varchar(200) NOT NULL,
  `Country` varchar(50) NOT NULL
);
6. Agregar las restricciones:
ALTER TABLE `usuarios`
  ADD PRIMARY KEY (`Email`),
  ADD UNIQUE KEY `Username` (`Username`);
7. Corra en su IDE de preferencial el main:
go run main.go
8. Abra la pagina web:
http://localhost:8080/

## Video de Presentacion de Idea 📢
https://youtu.be/2cARZhfOWv4

