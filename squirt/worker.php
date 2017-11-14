<?php

`touch /tmp/huh.txt`;
while($f = fgets(STDIN)){
    echo "line: $f";
    `touch /tmp/bleh.txt`;
}