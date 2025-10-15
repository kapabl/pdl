<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\FieldList;

class SessionFieldList extends FieldList
{
    protected function setup()
    {
        $this->addField( 'id' );
        $this->addField( 'ip_address' );
        $this->addField( 'last_activity' );
        $this->addField( 'payload' );
        $this->addField( 'user_agent' );
        $this->addField( 'user_id' );
    }
}

