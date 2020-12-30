<?php // Do not put any HTML above this line
    session_start();
    require_once "bootstrap.php";
    require_once "pdo.php";
if ( isset($_POST['cancel'] ) ) {
    // Redirect the browser to game.php
    header("Location: index.php");
    return;
}



$salt = 'XyZzy12*_';
$stored_hash = '1a52e17fa899cf40fb04cfc42e6352f1';  // Pw is php123

    // Check to see if we have some POST data, if we do process it
    if ( isset($_POST['email']) && isset($_POST['pass']) ) {
        if (stripos($_POST['email'], '@')== false) {
                $_SESSION['error'] = "Email must have an at-sign (@)";
                header("Location: login.php");
                return;
                
                if ( strlen($_POST['email']) < 1 || strlen($_POST['pass']) < 1 ){
                   $_SESSION['error'] = "Email and password are required";
                    header("Location: login.php");
                    return;
                }
            }else {
                $check = hash('md5', $salt.$_POST['pass']);
                if ( $check == $stored_hash ) {
                    // Redirect the browser to index.php
                    $_SESSION['name'] = $_POST['email'];
                    header("Location: index.php");
                    error_log("Login success ".$_POST['email']);
                    error_log("Login fail ".$_POST['email']." $check");
                    return;
                } else {
                    $_SESSION['error'] = "Incorrect password";
                    header("Location: login.php");
                    return;    
            }
        }
    }

// Fall through into the View
?>
<!DOCTYPE html>
<html>
    <head>
        <?php require_once "bootstrap.php"; ?>
        <title>Ahmad Muhammad Ateya 6f46dd17</title>
    </head>
    <body>
        <div class="container">
            <h1>Please Log In</h1>

                <?php
                if ( isset($_SESSION['error']) ) {
                    echo('<p style="color: red;">'.htmlentities($_SESSION['error'])."</p>\n");
                    unset($_SESSION['error']);
                    }
                ?>

            <form method="POST" action="login.php">
                <label for="nam">Email</label>
                <input type="text" name="email" id="nam">
                <br/>
                
                <label for="id_1723">Password</label>
                <input type="password" name="pass" id="id_1723">
                <br/>

                <input type="submit" value="Log In">
                <input type="submit" name="cancel" value="Cancel">
            </form>

            <p>
            For a password hint, view source and find an account and password hint
            in the HTML comments.
            <!-- Hint:
            The account is csev@umich.edu
            The password is the three character name of the
            programming language used in this class (all lower case)
            followed by 123. -->
            </p>
        </div>
    </body>
</html>