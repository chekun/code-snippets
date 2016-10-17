<?php

//图片版权(http://weibo.com/1751035982/E9SV40tEU?from=page_1003061751035982_profile&wvr=6&mod=weibotime&type=comment#_rnd1476673547887)
$im = new Imagick('./hebe.jpg');

$layer = new Imagick();
$layer->newImage($im->getImageWidth(), $im->getImageHeight(), '#C0FFFF');
$layer->setImageOpacity (0.44);
$im->compositeImage($layer, Imagick::COMPOSITE_SOFTLIGHT, 0, 0);

$layer = new Imagick();
$layer->newImage($im->getImageWidth(), $im->getImageHeight(), '#000699');
$layer->setImageOpacity (0.48);
$im->compositeImage($layer, Imagick::COMPOSITE_EXCLUSION, 0, 0);

$im->writeImage('./vintage.jpg');