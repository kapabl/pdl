<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\OrderBy;

class MigrationOrderBy extends OrderBy
{
    protected function setup()
    {
        $this->addField( 'batch' );
        $this->addField( 'id' );
        $this->addField( 'migration' );
    }
}

