<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\Where;

class CacheWhere extends Where
{
    protected function setup()
    {
        $this->addField( 'expiration' );
        $this->addField( 'key' );
        $this->addField( 'value' );
    }
}

