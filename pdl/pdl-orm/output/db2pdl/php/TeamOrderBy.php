<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\OrderBy;

class TeamOrderBy extends OrderBy
{
    protected function setup()
    {
        $this->addField( 'created_at' );
        $this->addField( 'id' );
        $this->addField( 'name' );
        $this->addField( 'personal_team' );
        $this->addField( 'updated_at' );
        $this->addField( 'user_id' );
    }
}

