<?php
include 'db_connect_parameters.php';
	try {
		if ($_SERVER['REQUEST_METHOD'] !== 'GET') {
			throw new Exception("Некорректный метод запроса");
		}
		//Подключение к БД
		$conn = new mysqli($servername, $username, $password, $dbname);
		//Проверка соединения
		if ($conn->connect_error) {
		    throw new Exception("Ошибка соединения с БД: " . $conn->connect_error);
		}
		//запрос к БД
		$sql = "SELECT id, dir_path, total_size, loading_time, date_of_request, time_of_request FROM stat";
		//результат запроса
		$result = $conn->query($sql);
		//данные для графика
		$data[] = array();
		//заполнение таблицы
		$tableBody = "";
		$tableBody. "<tr>";
		if ($result->num_rows > 0) {
		    while($row = $result->fetch_assoc()) {
		    	if (isset($row["id"]) && isset($row["dir_path"]) && isset($row["total_size"]) && isset($row["loading_time"]) && isset($row["date_of_request"]) && isset($row["time_of_request"])){
				$tableBody. "<tr>".
				$tableBody. "<td>" . $row["id"] . "</td>".
				$tableBody. "<td class = 'left'>" . $row["dir_path"] . "</td>".
				$tableBody. "<td class = 'left'>" . $row["total_size"] . "</td>".
				$tableBody. "<td class = 'center'>" . $row["loading_time"] . "</td>".
				$tableBody. "<td class = 'center'>" . $row["date_of_request"] . "</td>".
				$tableBody. "<td class = 'center'>" . $row["time_of_request"] . "</td>".
				$tableBody. "</tr>";
		    		
		    		array_push($data, [$row["total_size"], $row["loading_time"]]);}
		    	else {throw new Exception("Отстутсвуют некоторые столбцы");}
	    		
	    		}
		       }
		 
		
		}
	catch(Exception $e){
		http_response_code(500);
		die("500 Internal Server Error: ". $e->getMessage(). "\n");
	}
	finally{
	//закрытие соединенния
		$conn->close();
	}
include 'template.html';
?>
