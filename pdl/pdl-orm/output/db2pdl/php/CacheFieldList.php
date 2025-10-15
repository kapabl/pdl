<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\FieldList;

class CacheFieldList extends FieldList
{
    protected function setup()
    {
        $this->addField( 'expiration' );
        $this->addField( 'key' );
        $this->addField( 'value' );
    }
}

