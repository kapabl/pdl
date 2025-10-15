<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\FieldList;

class MigrationFieldList extends FieldList
{
    protected function setup()
    {
        $this->addField( 'batch' );
        $this->addField( 'id' );
        $this->addField( 'migration' );
    }
}

