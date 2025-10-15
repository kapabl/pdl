<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\Where;

class FailedJobWhere extends Where
{
    protected function setup()
    {
        $this->addField( 'connection' );
        $this->addField( 'exception' );
        $this->addField( 'failed_at' );
        $this->addField( 'id' );
        $this->addField( 'payload' );
        $this->addField( 'queue' );
        $this->addField( 'uuid' );
    }
}

