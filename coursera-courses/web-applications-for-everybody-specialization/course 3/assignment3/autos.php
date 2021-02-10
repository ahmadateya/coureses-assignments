<?php
include 'pdo.php';
// Demand a GET parameter
if ( ! isset($_GET['name']) || strlen($_GET['name']) < 1  ) {
    die('Name parameter missing');
}

// If the user requested logout go back to index.php
if ( isset($_POST['logout']) ) {
    header("Location: index.php");
    return;
}

    $username = $_GET['name'];
    $failure = false;
    $success = false;

// Check to see if we have some POST data, if we do process it
if ( isset($_POST['add'])) {
    if ( is_numeric($_POST['year']) && is_numeric($_POST['mileage']) ) 
    {
           if (strlen($_POST['make']) < 1) 
           {
               $failure = "Make is required";
           }else{
                $stmt= $pdo->prepare("INSERT INTO `autos`(make, year, mileage) 
                                        VALUES (:make, :year, :mileage)");
                $stmt->execute(array(
                    ':make'     => $_POST['make'],
                    ':year'     => $_POST['year'],
                    ':mileage'  => $_POST['mileage'] ));

                $success = "Record inserted";
           }
    }else
        {
            $failure = "Mileage and year must be numeric";
        }
}
?>

<!DOCTYPE html>
<html>
<head>
<title>Ahmad Muhammad Ateya 6f46dd17</title>
<?php require_once "bootstrap.php"; ?>
</head>
<body>
    <div class="container">
        <h1>Tracking Autos for <?php echo $username; ?></h1>


            <?php
                if ( $failure !== false ) {
                    echo("<p style=\"color: red;\">".htmlentities($failure)."</p>\n");
                }
            ?>

            <?php
                if ( $success !== false ) {
                    echo("<p style=\"color: green;\">".htmlentities($success)."</p>\n");
                }
            ?>

        <form method="POST" action="autos.php <?php echo "?name=".urlencode($_GET['name']) ?>">
            <label>make:</label>
            <input type="text" name="make">
            <br>

            <label>Year:</label>
            <input type="text" name="year">
            <br>

            <label>Mileage:</label>
            <input type="text" name="mileage">
            <br>
            <input type="submit" name="add" value="Add">
            <input type="submit" name="logout" value="Logout">
        </form>

        <h2>Automobiles</h2>
        <!-- Fetching the data from db -->
        <?php $stmt = $pdo->query("SELECT * FROM  `autos`"); ?>
                <ul>
                <?php while ($row = $stmt->fetch(PDO::FETCH_ASSOC)): ?>
                    <li><?php echo (htmlentities($row['year'])." ".
                     htmlentities($row['make'])."/".
                     htmlentities($row['mileage'])); ?>
                    </li>
                <?php endwhile ?>
                 </ul>  
    </div>
</body>
</html>
