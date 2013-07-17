<?php
/** @var Base $f3 */
$f3=require('lib/base.php');

$f3->set('DEBUG',0);
$f3->set('UI','ui/');

// https://github.com/TechEmpower/FrameworkBenchmarks#json-response
$f3->route('GET /json',function($f3) {
    /** @var Base $f3 */
    header("Content-type: application/json");
    return $f3->serialize(array('message' => 'Hello World!'));
});


// https://github.com/TechEmpower/FrameworkBenchmarks#database-single-query
// https://github.com/TechEmpower/FrameworkBenchmarks#database-multiple-queries
$f3->route(
    array(
        'GET /db',
        'GET /db/@queries',
    ),
    function ($f3,$params) {
    /** @var Base $f3 */
        $params += array('queries' => 1); //default value
        $db = new \DB\SQL('mysql:host=localhost;port=3306;dbname=hello_world',
                          'benchmarkdbuser', 'benchmarkdbpass');
        $result = array();
        for ($i = 0; $i < $params['queries']; ++$i) {
            $id = mt_rand(1, 10000);
            $result[] = $db->exec('SELECT randomNumber FROM World WHERE id = ?',$id,0,false);
        }

        header("Content-type: application/json");
        return $f3->serialize($result);
    }
);

// https://github.com/TechEmpower/FrameworkBenchmarks#database-single-query
// https://github.com/TechEmpower/FrameworkBenchmarks#database-multiple-queries
$f3->route(
    array(
         'GET /db-orm',
         'GET /db-orm/@queries',
    ),
    function ($f3, $params) {
        /** @var Base $f3 */
        $params += array('queries' => 1); //default value
        $db = new \DB\SQL('mysql:host=localhost;port=3306;dbname=hello_world',
            'benchmarkdbuser', 'benchmarkdbpass');
        $mapper = new \DB\SQL\Mapper($db,'World');
        $result = array();
        for ($i = 0; $i < $params['queries']; ++$i) {
            $id = mt_rand(1, 10000);
            $result[] = $mapper->findone(array('where id = ?',$id))->cast();
        }

        header("Content-type: application/json");
        return $f3->serialize($result);
    }
);


$f3->route('GET /plaintext', function ($f3) {
    echo "Hello, World!";
});


$f3->route('GET /fortune', function ($f3) {
    /** @var Base $f3 */

    // to be continued
});


$f3->route(
    array(
         'GET /updateraw',
         'GET /updateraw/@queries',
    ),function($f3,$params) {
    /** @var Base $f3 */
    $params += array('queries' => 1); //default value

    // to be continued
});

$f3->run();
