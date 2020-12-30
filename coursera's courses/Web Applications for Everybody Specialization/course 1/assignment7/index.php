<!DOCTYPE html>
<head><title>Ahmad Muhammad Ateya MD5 Cracker</title></head>
<body>
<h1>MD5 cracker</h1>
<p>This application takes an MD5 hash
of four digits and check all 10.000 possible four digits PINs
to determine the PIN.</p>
<pre>
Debug Output:
<?php
$ans = "Not found";
// If there is no parameter, this code is all skipped
if ( isset($_GET['md5']) ) {
    $time_pre = microtime(true);
    $md5 = $_GET['md5'];

    // This is our numbers
    $txt = "0123456789";
    $show = 15;

    
    for($i=0; $i<strlen($txt); $i++ ) {
        $num1 = $txt[$i];   // The first num

        
        for($j=0; $j<strlen($txt); $j++ ) {
            $num2 = $txt[$j];  // Our second num


            for($k=0; $k<strlen($txt); $k++ ) {
                $num3 = $txt[$k];  // Our third num


                for($l=0; $l<strlen($txt); $l++ ) {
                    $num4 = $txt[$l];  // Our fourth num


                        $try = $num1.$num2.$num3.$num4;

                        // Run the hash and then check to see if we match
                        $check = hash('md5', $try);
                        if ( $check == $md5 ) {
                            $ans = $try;
                            break;   
                        }

                        // Debug output until $show hits 0
                        if ( $show > 0 ) {
                            print "$check $try\n";
                            $show = $show - 1;
                        }
                       
                }
            }
        }
    }
    // Compute elapsed time
    $time_post = microtime(true);
    print "Elapsed time: ";
    print $time_post-$time_pre;
    print "\n";

}
?>
</pre>
<!-- Use the very short syntax and call htmlentities() -->
<p>PIN: <?php echo $ans; ?></p>
<form>
<input type="text" name="md5" size="60" />
<input type="submit" value="Crack MD5"/>
</form>
<ul>
<li><a href="index.php">Reset</a></li>
<li><a href="md5.php">MD5 Encoder</a></li>
<li><a href="makecode.php">MD5 Code Maker</a></li>
</ul>
</body>
</html>

