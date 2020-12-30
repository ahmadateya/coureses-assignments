<?php 
	require_once "bootstrap.php";
 	require_once "pdo.php";
 	session_start();
  ?>

<!DOCTYPE html>
<html>
<head>
		<title>Ahmad Muhammad Ateya 6f46dd17</title>
</head>
<body>
	<div class="container">
		<?php
			echo('<h2>Welcome to the Automobiles Database</h2>'."\n");

			if (isset($_SESSION['name'])) {
				if ( isset($_SESSION['error']) ) {
			    echo '<p style="color:red">'.$_SESSION['error']."</p>\n";
			    unset($_SESSION['error']);
				}

				if ( isset($_SESSION['success']) ) {
				    echo '<p style="color:green">'.$_SESSION['success']."</p>\n";
				    unset($_SESSION['success']);
				}
					
					echo('<table border="2">'."\n");
					$stmt = $pdo->query("SELECT autos_id, make, model, year, mileage FROM autos2");
					
					echo ("<thead><tr>");
					echo ("<th>Make</th>");
					echo ("<th>Model</th>");
					echo ("<th>Year</th>");
					echo ("<th>Mileage</th>");
					echo ("<th>Action</th>");
					echo ("</tr></thead>");

				 	while ($row = $stmt->fetch(PDO::FETCH_ASSOC) ) {
					    echo "<tr><td>";
					    	echo(htmlentities($row['make']));
					    echo("</td><td>");
					    	echo(htmlentities($row['model']));
					    echo("</td><td>");
					    	echo(htmlentities($row['year']));
					    echo("</td><td>");
					    	echo(htmlentities($row['mileage']));
				    	echo("</td><td>");
					    echo('<a href="edit.php?autos_id='.$row['autos_id'].'">Edit</a> / ');
					    echo('<a href="delete.php?autos_id='.$row['autos_id'].'">Delete</a>');
					    echo("</td></tr>\n");
			 	}
				
				echo('</table>'."\n");
				echo('<a href="add.php">Add New Entry</a>'."\n");
			}else{
				echo('<p><a href="login.php">Please log in</a></p>'."\n");
				echo('<p>Attempt to <a href="add.php">add data</a> without logging in</p>'."\n");
			}
		?>
	</div>
</body>
</html>


