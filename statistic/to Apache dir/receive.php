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
	if (isset($data['full_size']) && isset($data['loading_time']) && isset($data['date_of_request']) && isset($data['time_of_request']) && isset($data['root'])){
		if (json_last_error() === JSON_ERROR_NONE) {
		    $full_size = $data['full_size'];
		    $loading_time = $data['loading_time'];
		    $date_of_request = $data['date_of_request'];
		    $time_of_request = $data['time_of_request'];
		    $dir_path = $data['root'];
		    //Вставка данных
		    $stmt = $conn->prepare("INSERT INTO stat (dir_path, total_size, loading_time, time_of_request, date_of_request) VALUES (?, ?, ?, ?, ?)");
		    $stmt->bind_param("sssss", $dir_path, $full_size,$loading_time,$time_of_request, $date_of_request);

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
