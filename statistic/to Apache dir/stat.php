<!DOCTYPE html>
<html>
    <!--Заголовок и подключение css-->
    <head>
        <title>Statistic</title>
        <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
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
    <button type="button" id = "buttonBack" onclick = "window.location.href='/'">Назад</button>
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
		$sql = "SELECT id, dir_path, total_size, loading_time, date_of_request, time_of_request FROM Stat";
		//результат запроса
		$result = $conn->query($sql);
		//заполнение таблицы
		$data[] = array();
		if ($result->num_rows > 0) {
		    while($row = $result->fetch_assoc()) {
		    	if (isset($row["id"]) && isset($row["dir_path"]) && isset($row["total_size"]) && isset($row["loading_time"]) && isset($row["date_of_request"]) && isset($row["time_of_request"])){
				echo "<tr>";
				echo "<td>" . $row["id"] . "</td>";
				echo "<td class = 'left'>" . $row["dir_path"] . "</td>";
				echo "<td class = 'left'>" . $row["total_size"] . "</td>";
				echo "<td class = 'center'>" . $row["loading_time"] . "</td>";
				echo "<td class = 'center'>" . $row["date_of_request"] . "</td>";
				echo "<td class = 'center'>" . $row["time_of_request"] . "</td>";
				echo "</tr>";
		    		
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
	?>
	</tbody>
</table>
<div id = "mistakeMessage"></div>
<canvas id="myChart" width="400" height="200"></canvas>
<script type = "module">
const ctx = document.getElementById('myChart');
const rawData = <?php echo json_encode($data)?>;
//сортировка данных для графика
rawData.sort((a, b) => a[0] - b[0]);

//график
const data = {
  datasets: [{
    label: 'График',
    data: rawData,
    backgroundColor: 'rgb(255, 99, 132)'
  }]
};
  new Chart(ctx, {
    type: 'line',
    data: data,
    order:0,
    options: {
    scales: {
      y: {
		beginAtZero: true,
      	type: 'logarithmic',
        title: {
          display: true,
          align: 'center',
          text: 'Время (миллисекунды)',
          color: 'black',
          font: {
            family: 'Arial',
            size: 14,
            weight: 'bold',
          }}},
          
      x: {
		beginAtZero: true,
      	type: 'logarithmic',
        title: {
          display: true,
          align: 'center',
          text: 'Размер (байты)',
          color: 'black',
          font: {
            family: 'Arial',
            size: 14,
            weight: 'bold',
          }}}}
    }
  });
</script>
</body>
</html>
