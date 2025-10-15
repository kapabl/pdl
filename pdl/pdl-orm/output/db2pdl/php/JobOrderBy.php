<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\OrderBy;

class JobOrderBy extends OrderBy
{
    protected function setup()
    {
        $this->addField( 'attempts' );
        $this->addField( 'available_at' );
        $this->addField( 'created_at' );
        $this->addField( 'id' );
        $this->addField( 'payload' );
        $this->addField( 'queue' );
        $this->addField( 'reserved_at' );
    }
}

