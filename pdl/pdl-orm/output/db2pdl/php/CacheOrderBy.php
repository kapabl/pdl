<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\OrderBy;

class CacheOrderBy extends OrderBy
{
    protected function setup()
    {
        $this->addField( 'expiration' );
        $this->addField( 'key' );
        $this->addField( 'value' );
    }
}

