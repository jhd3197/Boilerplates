<?php

if (!defined('ABSPATH')) {
    exit;
}

class MyProject_Activator {
    public static function activate() {
        if (!get_option('myproject_version')) {
            add_option('myproject_version', '0.1.0');
        }
    }
}
