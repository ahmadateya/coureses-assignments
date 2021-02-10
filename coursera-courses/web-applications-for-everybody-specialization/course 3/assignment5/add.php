<?php 
    session_start();
    include 'pdo.php';

    if ( ! isset($_SESSION['name']) ) {
        die('ACCESS DENIED');
    }

    if (isset($_POST['logout'])) {
       header('Location: logout.php');
    }


// Check to see if we have some POST data, if we do process it
    if ( isset($_POST['add'])) {

        if ( strlen($_POST['make']) < 1 || strlen($_POST['model']) < 1|| strlen($_POST['year']) < 1 || strlen($_POST['make']) < 1 ) {

                $_SESSION['error'] = "All values are required";
                           header("Location: add.php");
                            return;

        }elseif( is_numeric($_POST['year']) && is_numeric($_POST['mileage']) ){

                $stmt= $pdo->prepare("INSERT INTO `autos2`(make, model, year, mileage) 
                                            VALUES (:make,:model, :year, :mileage)");
                    $stmt->execute(array(
                        ':make'     => $_POST['make'],
                        ':model'     => $_POST['model'],
                        ':year'     => $_POST['year'],
                        ':mileage'  => $_POST['mileage'] ));

                    $_SESSION['success'] = "Record added.";
                    header("Location: index.php");
                    return;

        }else{
            $_SESSION['error'] = "Mileage and year must be numeric";
                    header("Location: add.php");
                    return;
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
        <h1>Tracking Autos for <?php $_SESSION['name'] ?></h1>

        <?php
            if ( isset($_SESSION['error']) ) {
                echo('<p style="color: red;">'.htmlentities($_SESSION['error'])."</p>\n");
                unset($_SESSION['error']);
                }
        ?>

        <form method="POST" action="add.php">
            <label>make:</label>
            <input type="text" name="make">
            <br>

            <label>model:</label>
            <input type="text" name="model">
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
    </div>
</body>
</html>