<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class WebsocketsStatisticsEntryColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'api_message_count', 'int' );
        $this->addColumn( 'app_id', 'string' );
        $this->addColumn( 'created_at', 'string' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'peak_connection_count', 'int' );
        $this->addColumn( 'updated_at', 'string' );
        $this->addColumn( 'websocket_message_count', 'int' );
    }
}

