<?php


main();

function main() {
	$input = loadInput("input.txt");

	Part1($input);
	echo "\n";
	Part2($input);
}

function loadInput(string $fileName) {
	return explode("\n", file_get_contents($fileName));
}

function Part1(array $input = []) {
	$sum = 0;

	foreach ($input as $v) {
		$sum += $v;
	}

	echo "Part 1\n";
	echo "\tSum: ", $sum;
}

function Part2(array $input = []) {
	$visited = [];

	$sum = 0;
	$firstDouble = 0;
	$boundary = 10000;
	$found = false;

	for ($i = 0; !$found && $i < $boundary; $i++) {
		foreach ($input as $_ => $v) {
			$sum += $v;

			if (isset($visited[$sum])) {
				$firstDouble = $sum;
				$found = true;
				break;
			}

			$visited[$sum] = true;
		}
	}

	if (!$found) {
		echo "Error occured, no doubles.\n";
		exit(1);
	}

	echo "Part 2:\n";
	echo "\tFirst Double: ", $firstDouble, "\n";
}