<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\Where;

class WebsocketsStatisticsEntryWhere extends Where
{
    protected function setup()
    {
        $this->addField( 'api_message_count' );
        $this->addField( 'app_id' );
        $this->addField( 'created_at' );
        $this->addField( 'id' );
        $this->addField( 'peak_connection_count' );
        $this->addField( 'updated_at' );
        $this->addField( 'websocket_message_count' );
    }
}

