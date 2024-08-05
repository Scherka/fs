<?php
include 'db_connect_parameters.php';
header('Content-Type: application/json');
try {
	if ($_SERVER['REQUEST_METHOD'] !== 'POST') {
		throw new Exception("Некорректный метод запроса");
	}
	//Подключение к БД
	$conn = new mysqli($servername, $username, $password, $dbname);
	//Проверка подключения
	if ($conn->connect_error) {
	    throw new Exception("Ошибка соединения c БД: " . $conn->connect_error);
	}
	//Получаем данные из JSON

	$json = file_get_contents('php://input');
	$data = json_decode($json, true);
	if (isset($data['FullSize']) && isset($data['LoadingTime']) && isset($data['DateOfRequest']) && isset($data['TimeOfRequest']) && isset($data['Root'])){
		if (json_last_error() === JSON_ERROR_NONE) {
		    $full_size = $data['FullSize'];
		    $loading_time = $data['LoadingTime'];
		    $DateOfRequest = $data['DateOfRequest'];
		    $TimeOfRequest = $data['TimeOfRequest'];
		    $dir_path = $data['Root'];
		    //Вставка данных
		    $stmt = $conn->prepare("INSERT INTO Stat (dir_path, total_size, loading_time, time_of_request, date_of_request) VALUES (?, ?, ?, ?, ?)");
		    $stmt->bind_param("sssss", $dir_path, $full_size,$loading_time,$TimeOfRequest, $DateOfRequest);

		    if ($stmt->execute()) {
		    	echo "Запись для $dir_path создана \n";
		    } else {
		    	throw new Exception("Ошибка при выполнении вставки: ". $stmt->error);
		    }

		    
		} else {
		    throw new Exception("Ошибка при чтении данных из JSON");
		}
	} else {
		throw new Exception("В JSON отсутствуют некоторые поля.");
	}
}
catch(Exception $e){
	http_response_code(500);
	die( "Ошибка: ". $e->getMessage(). "\n");
}
finally{
$stmt->close();
$conn->close();
}
?>
