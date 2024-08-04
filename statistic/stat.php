<?php
$servername = "localhost";
$username = "root";
$password = "12345678";
$dbname = "statDB";
// Create connection
$conn = new mysqli($servername, $username, $password, $dbname);
// Check connection
if ($conn->connect_error) {
    die("Ошибка соединения с БД: " . $conn->connect_error);
}
$sql = "SELECT id, dir_path, total_size, loading_time, date_of_request, time_of_request FROM Stat";
$result = $conn->query($sql);
?>
<!DOCTYPE html>
<html>
    <!--Заголовок и подключение css-->
    <head>
        <title>Statistic</title>
        <meta charset=utf-8>
        <style>
        .table {
	/*Стиль таблицы */
	width: 100%;
	margin-bottom: 20px;
	border: 1px solid #dddddd;
	border-collapse: collapse; 

	}
	/*Стиль заголовка */
	.table th {
		font-weight: bold;
		padding: 5px;
		background: #ebebeb;
		border: 1px solid #dddddd;
	}
	/*Стиль ячеек */
	.table td {
		border: 1px solid #dddddd;
		padding: 5px;
	}
	.left {
		text-align: left;
	}
	.center {
		text-align: center;
	}
        </style>
    </head>  
<body>
    <!--Кнопки-->
    <div class = "layer1">
    </div>
<!--таблица-->
<table class="table">
	
    <caption class = "layer1">Запросы</div></caption>
	<thead>  
		<tr>
			<th>id</th>
			<th>Путь</th>
			<th>Размер</th>
			<th>Время выполнения запроса</th>
			<th>Дата запроса</th>
			<th>Время запроса</th>
		</tr>
	</thead>
	<tbody>
	<?php if ($result->num_rows > 0) {
	    while($row = $result->fetch_assoc()) {
		echo "<tr>";
		echo "<td>" . $row["id"] . "</td>";
		echo "<td class = 'path'>" . $row["dir_path"] . "</td>";
		echo "<td class = 'left'>" . $row["total_size"] . "</td>";
		echo "<td class = 'center'>" . $row["loading_time"] . "</td>";
		echo "<td class = 'center'>" . $row["date_of_request"] . "</td>";
		echo "<td class = 'center'>" . $row["time_of_request"] . "</td>";
		echo "</tr>";
    		}
	       } 
	$conn->close();
	?>
	</tbody>
</table>
</body>
</html>
