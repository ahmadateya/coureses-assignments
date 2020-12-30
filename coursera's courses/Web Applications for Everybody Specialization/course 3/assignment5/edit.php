<?php
require_once "pdo.php";
require_once "bootstrap.php";
session_start();


    if ( isset($_POST['submit'])) {

        if ( strlen($_POST['make']) < 1 || strlen($_POST['model']) < 1|| strlen($_POST['year']) < 1 || strlen($_POST['make']) < 1 ) {

                $_SESSION['error'] = "All values are required";
                           header("Location: edit.php?autos_id=".$_POST['autos_id']);
                            return;

        }elseif( is_numeric($_POST['year']) && is_numeric($_POST['mileage']) ){

                $sql = "UPDATE autos2 SET make = :make,
                        model = :model, year = :year, mileage=:mileage
                        WHERE autos_id = :autos_id";


                $stmt = $pdo->prepare($sql);
                $stmt->execute(array(
                    ':make' => $_POST['make'],
                    ':model' => $_POST['model'],
                    ':year' => $_POST['year'],
                    ':mileage' => $_POST['mileage'],
                    ':autos_id' => $_POST['autos_id']));
                $_SESSION['success'] = 'Record updated';
                header( 'Location: index.php' );
                return;

        }else{
            $_SESSION['error'] = "Mileage and year must be numeric";
                    header("Location: edit.php?autos_id=".$_POST['autos_id']);
                    return;
        }
                         
    }








// Guardian: Make sure that autos_id is present
if ( ! isset($_GET['autos_id']) ) {
  $_SESSION['error'] = "Missing autos_id";
  header('Location: index.php');
  return;
}


        $stmt = $pdo->prepare("SELECT * FROM autos2 where autos_id = :xyz");
        $stmt->execute(array(":xyz" => $_GET['autos_id']));
        $row = $stmt->fetch(PDO::FETCH_ASSOC);
        if ( $row === false ) {
            $_SESSION['error'] = 'Bad value for autos_id';
            header( 'Location: index.php');
            return;
        }



$make = htmlentities($row['make']);
$model = htmlentities($row['model']);
$year = htmlentities($row['year']);
$mileage = htmlentities($row['mileage']);

?>
<!DOCTYPE html>
<html>
<head>
        <title>Ahmad Muhammad Ateya 6f46dd17</title>
</head>
<body>
<div class="container">

        
    <h2>Editing Automobile</h2>

    <?php
            // Flash message
        if ( isset($_SESSION['error']) ) {
            echo '<p style="color:red">'.$_SESSION['error']."</p>\n";
            unset($_SESSION['error']);
        }   
        ?>
        
    <form method="post" action="edit.php?autos_id=' <?. $_GET['autos_id']; ?>">
    <p>make:
    <input type="text" name="make" value="<?= $make ?>"></p>
    <p>model:
    <input type="text" name="model" value="<?= $model ?>"></p>
    <p>year:
    <input type="text" name="year" value="<?= $year ?>"></p>
    <p>mileage:
    <input type="text" name="mileage" value="<?= $mileage ?>"></p>
    <input type="hidden" name="autos_id" value="<?= $_GET['autos_id']; ?>">
    <p><input type="submit" name="submit" value="Save"/>
    <a href="index.php">Cancel</a></p>
    </form>
</div>
</body>
</html>
