package linttest_test

import (
	"testing"

	"github.com/VKCOM/noverify/src/linttest"
)

func TestTraitProperties(t *testing.T) {
	linttest.SimpleNegativeTest(t, `<?php
	declare(strict_types=1);

	trait Example
	{
		private static $property = 'some';

		protected function some(): string
		{
			return self::$property;
		}
	}`)
}

func TestTraitSelf(t *testing.T) {
	linttest.SimpleNegativeTest(t, `<?php
define('null', 0);

function define($name, $v) {}

trait SingletonSelf {
    /** @var self */
    private static $instance = null;

    /** @return self */
    public static function instance() {
        if (self::$instance === null) {
            self::$instance = new self();
        }

        return self::$instance;
    }
}

trait SingletonStatic {
    /** @var static */
    private static $instance = null;

    /** @return static */
    public static function instance() {
        if (static::$instance === null) {
            static::$instance = new static();
        }

        return static::$instance;
    }
}
`)
}
