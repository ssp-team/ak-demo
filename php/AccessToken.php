<?php

require_once("xxtea.php");

function getMillisecond(){
	list($s1,$s2)=explode(' ',microtime());
	return (float)sprintf('%.0f',(floatval($s1)+floatval($s2))*1000);
}


function createAccessToken($accessKeyId, $mediaId, $accessKeySecret){
	$ts=getMillisecond();
	$str= $accessKeyId .'|' . $ts . '|' . $mediaId;
	$encrypt_data = xxtea_encrypt($str, $accessKeySecret);
	# $decrypt_data = xxtea_decrypt($encrypt_data, $accessKeySecret);
	$b64 = base64_encode($encrypt_data);
	$b64 = str_replace('+','-',$b64);
	$b64 = str_replace('/','_',$b64);
	$b64 = str_replace('=','*',$b64);
	return $b64;
}

$accessKeyId = "[accessKeyId Here]";
$accessKeySecret = '[accessKeySecret Here]';
$mediaId="[mediaId Here]";

$b64 = createAccessToken($accessKeyId,$mediaId,$accessKeySecret);

echo 'AccessToken: '. PHP_EOL . $b64. PHP_EOL;



