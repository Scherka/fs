<?php
$servername = "localhost";
$username = "root";
$password = "12345678";
$dbname = "statDB";
header('Content-Type: application/json');
// Create connection
$conn = new mysqli($servername, $username, $password, $dbname);
// Check connection
if ($conn->connect_error) {
    die("Connection failed: " . $conn->connect_error);
}
// Get the JSON data from the POST request
$json = file_get_contents('php://input');
$data = json_decode($json, true);
if (json_last_error() === JSON_ERROR_NONE) {
    $full_size = $data['FullSize'];
    $loading_time = $data['LoadingTime'];
    $DateOfRequest = $data['DateOfRequest'];
    $TimeOfRequest = $data['TimeOfRequest'];
    $dir_path = $data['Root'];
    // Insert data into the database
    $stmt = $conn->prepare("INSERT INTO Stat (dir_path, total_size, loading_time, time_of_request, date_of_request) VALUES (?, ?, ?, ?, ?)");
    $stmt->bind_param("sssss", $dir_path, $full_size,$loading_time,$TimeOfRequest, $DateOfRequest);

    if ($stmt->execute()) {
        echo "New record created successfully";
    } else {
        echo "Error: " . $stmt->error;
    }

    $stmt->close();
} else {
    echo "Invalid JSON data";
}

$conn->close();

?>