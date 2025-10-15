<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\Where;

class MigrationWhere extends Where
{
    protected function setup()
    {
        $this->addField( 'batch' );
        $this->addField( 'id' );
        $this->addField( 'migration' );
    }
}

