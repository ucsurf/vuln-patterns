<?php
function a($xml) {
  // ruleid: svg-doctype-regex-strip
  $xml = preg_replace('/<!DOCTYPE[^>]*>/i', '', $xml);
  // ruleid: svg-doctype-regex-strip
  $xml = str_ireplace('<!ENTITY', '', $xml);
  // ok: svg-doctype-regex-strip
  $xml = preg_replace('/<script/i', '', $xml);
  return $xml;
}
