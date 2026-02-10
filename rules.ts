type Severity = "error" | "warning" | "ignore" | "forbidden";

const ALL_OPTIONS = "all_options";

type RuleConfig = {
    severity: Severity;
}

type WithOptions<T> = T & {
    options: string[];
}

type WithCompare<T> = T & {
    compare: {key: string; value: string}[];
}

type GtfsRules = {
    agency: {
        _file: Severity;
        agency_id: WithOptions<RuleConfig>;
        agency_name: WithOptions<RuleConfig>;
        agency_name_id_match: WithCompare<RuleConfig>;
        agency_url: RuleConfig;
        agency_timezone: RuleConfig;
        agency_lang: RuleConfig;
        agency_phone: RuleConfig;
        agency_fare_url: RuleConfig;
        agency_email: RuleConfig;
    }
    stops: {
        _file: Severity;
        stop_id: RuleConfig;
        stop_code: RuleConfig;
        stop_name: RuleConfig;
        stop_short_name: RuleConfig;
        tts_stop_name: RuleConfig;
        stop_desc: RuleConfig;
        stop_lat: RuleConfig;
        stop_lon: RuleConfig;
        zone_id: RuleConfig;
        stop_url: RuleConfig;
        location_type: RuleConfig;
        parent_station: RuleConfig;
        stop_timezone: RuleConfig;
        wheelchair_boarding: WithOptions<RuleConfig>;
        level_id: RuleConfig;
        platform_code: RuleConfig;
        public_visible: WithOptions<RuleConfig>;
        has_stop_sign: WithOptions<RuleConfig>;
        has_shelter: WithOptions<RuleConfig>;
        shelter_code: RuleConfig;
        shelter_maintainer: RuleConfig;
        has_bench: WithOptions<RuleConfig>;
        has_network_map: WithOptions<RuleConfig>;
        has_schedules: WithOptions<RuleConfig>;
        has_pip_real_time: WithOptions<RuleConfig>;
        has_tariffs_information: WithOptions<RuleConfig>;
        region_id: WithOptions<RuleConfig>;
        municipality_id: WithOptions<RuleConfig>;
        parish_id: RuleConfig;
    }
    routes:{
        _file: Severity;
        line_id:RuleConfig;
        line_short_name:RuleConfig;
        line_long_name:RuleConfig;
        route_id:RuleConfig;
        agency_id:RuleConfig;
        route_short_name:RuleConfig;
        route_long_name:RuleConfig;
        route_desc:RuleConfig;
        route_remarks:RuleConfig;
        route_type:WithOptions<RuleConfig>;
        path_type:WithOptions<RuleConfig>;
        circular:WithOptions<RuleConfig>;
        school:WithOptions<RuleConfig>;
        route_url:RuleConfig;
        route_color:RuleConfig;
        route_text_color:RuleConfig;
        continuous_pickup:WithOptions<RuleConfig>;
        continuous_drop_off:WithOptions<RuleConfig>;
    }
    trips: {
        _file: Severity;
        route_id: RuleConfig;
        pattern_id: RuleConfig;
        service_id: RuleConfig;
        trip_id_limit_characters: RuleConfig;
        trip_id: RuleConfig;
        trip_headsign: RuleConfig;
        trip_short_name: RuleConfig;
        direction_id: RuleConfig;
        block_id: RuleConfig;
        shape_id: RuleConfig;
        wheelchair_accessible: RuleConfig;
        bikes_allowed: RuleConfig;
        stop_sequence: RuleConfig;
        direction_pattern_id_match: RuleConfig;
    }
    stop_times: {
        _file: Severity;
        trip_id: RuleConfig;
        arrival_time: RuleConfig;
        departure_time: RuleConfig;
        stop_id: RuleConfig;
        stop_sequence: RuleConfig;
        stop_headsign: RuleConfig;
        pickup_type: RuleConfig;
        drop_off_type: RuleConfig;
        continuous_pickup: RuleConfig;
        continuous_drop_off: RuleConfig;
        shape_dist_traveled: RuleConfig;
        start_pickup_drop_off_window: RuleConfig;
        end_pickup_drop_off_window: RuleConfig;
        pickup_booking_rule_id: RuleConfig;
        drop_off_booking_rule_id: RuleConfig;
        timepoint: RuleConfig;
        zone_1: RuleConfig;
        zone_2: RuleConfig;
        zone_3: RuleConfig;
    }
    calendar: {
        _file: Severity;
        service_id: RuleConfig;
        monday: RuleConfig;
        tuesday: RuleConfig;
        wednesday: RuleConfig;
        thursday: RuleConfig;
        friday: RuleConfig;
        saturday: RuleConfig;
        sunday: RuleConfig;
        start_date: RuleConfig;
        end_date: RuleConfig;
    }
    calendar_dates: {
        _file: Severity;
        service_id: RuleConfig;
        date: RuleConfig;
        exception_type: RuleConfig;
    }
    vehicles: {
        _file: Severity;
        vehicle_id: RuleConfig;
        agency_id: RuleConfig;
        license_plate: RuleConfig;
        make: RuleConfig;
        model: RuleConfig;
        owner: RuleConfig;
        registration_date: RuleConfig;
        available_seats: RuleConfig;
        available_standing: RuleConfig;
        typology: RuleConfig;
        propulsion: WithOptions<RuleConfig>;
        emission: WithOptions<RuleConfig>;
        climatization: WithOptions<RuleConfig>;
        wheelchair: WithOptions<RuleConfig>;
        lowered_floor: WithOptions<RuleConfig>;
        ramp: WithOptions<RuleConfig>;
        kneeling: WithOptions<RuleConfig>;
        static_information: WithOptions<RuleConfig>;
        onboard_monitor: WithOptions<RuleConfig>;
        front_display: WithOptions<RuleConfig>;
        rear_display: WithOptions<RuleConfig>;
        side_display: WithOptions<RuleConfig>;
        internal_sound: WithOptions<RuleConfig>;
        external_sound: WithOptions<RuleConfig>;
        consumption_meter: WithOptions<RuleConfig>;
        bicycles: WithOptions<RuleConfig>;
        passenger_counting: WithOptions<RuleConfig>;
        video_surveillance: WithOptions<RuleConfig>;
    }
    fare_attributes: {
        _file: Severity;
        fare_id: RuleConfig;
        price: RuleConfig;
        currency_type: RuleConfig;
        payment_method: WithOptions<RuleConfig>;
        transfers: WithOptions<RuleConfig>;
        agency_id: RuleConfig;
        transfer_duration: RuleConfig;
    }
    fare_rules: {
        _file: Severity;
        fare_id: RuleConfig;
        route_id: RuleConfig;
        origin_id: RuleConfig;
        destination_id: RuleConfig;
        contains_id: RuleConfig;
    }
    shapes: {
        _file: Severity;
        shape_id: RuleConfig;
        shape_pt_lat: RuleConfig;
        shape_pt_lon: RuleConfig;
        shape_pt_sequence: RuleConfig;
        shape_dist_traveled: RuleConfig;
    }
    frequencies: {
        _file: Severity;
        trip_id: RuleConfig;
        start_time: RuleConfig;
        end_time: RuleConfig;
        headway_secs: RuleConfig;
        exact_times: RuleConfig;
    }
    transfers: {
        _file: Severity;
        from_stop_id: RuleConfig;
        to_stop_id: RuleConfig;
        transfer_type: RuleConfig;
        min_transfer_time: RuleConfig;
    }
    pathways: {
        _file: Severity;
        pathway_id: RuleConfig;
        from_stop_id: RuleConfig;
        to_stop_id: RuleConfig;
        pathway_mode: RuleConfig;
        is_bidirectional: RuleConfig;
        length: RuleConfig;
        traversal_time: RuleConfig;
        stair_count: RuleConfig;
        max_slope: RuleConfig;
        min_width: RuleConfig;
        signposted_as: RuleConfig;
        reversed_signposted_as: RuleConfig;
    }
    levels: {
        _file: Severity;
        level_id: RuleConfig;
        level_index: RuleConfig;
        level_name: RuleConfig;
    }
    feed_info: {
        _file: Severity;
        feed_type: WithOptions<RuleConfig>;
        feed_publisher_name: RuleConfig;
        feed_publisher_url: RuleConfig;
        feed_lang: WithOptions<RuleConfig>;
        default_lang: RuleConfig;
        feed_start_date: RuleConfig;
        feed_end_date: RuleConfig;
        feed_version: RuleConfig;
        feed_remarks: RuleConfig;
        feed_contact_email: RuleConfig;
        feed_contact_url: RuleConfig;
    }
    translations: {
        _file: Severity;
        table_name: RuleConfig;
        field_name: RuleConfig;
        language: RuleConfig;
        translation: RuleConfig;
        record_id: RuleConfig;
        record_sub_id: RuleConfig;
        field_value: RuleConfig;
    }
    attributions: {
        _file: Severity;
        attribution_id: RuleConfig;
        agency_id: RuleConfig;
        route_id: RuleConfig;
        trip_id: RuleConfig;
        organization_name: RuleConfig;
        is_producer: RuleConfig;
        is_operator: RuleConfig;
        is_authority: RuleConfig;
        attribution_url: RuleConfig;
        attribution_email: RuleConfig;
        attribution_phone: RuleConfig;
    }
}

const rules: GtfsRules = {
    agency: {
        _file: "error",
        agency_id: {
            severity: "error",
            options: ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "18", "21", "23", "24", "34", "41", "42", "43", "44", "49", "54"]
        },
       "agency_name": {
            "severity": "error",
            "options": ["Área Metropolitana de Lisboa","Carris","Metropolitano Olissipo","Comboios de Portugal","Transtejo / Soflusa","Transportes Sul do Tejo","Rodoviária de Lisboa","Soflusa","Transportes Coletivos do Barreiro","Vimeca Transportes","Scotturb","ID/JJ/HLM","Isidoro Duarte","Barraqueiro Transportes","Joaquim Jerónimo","Fertagus","Metro Transportes do Sul","Henrique Leonardo da Mota","Cascais Próxima","Portal VIVA","Rodoviária do Tejo","Câmara Municipal de Lisboa","Viação Alvorada","Alsa Todi","Município de Oeiras","Municipio de Setúbal"]
        },
        "agency_name_id_match": {
            "severity": "error",
            "compare": [
                {"key": "0","value": "Área Metropolitana de Lisboa"},
                {"key": "1","value": "Carris"},
                {"key": "2","value": "Metropolitano Olissipo"},
                {"key": "3","value": "Comboios de Portugal"},
                {"key": "4","value": "Transtejo / Soflusa"},
                {"key": "5","value": "Transportes Sul do Tejo"},
                {"key": "6","value": "Rodoviária de Lisboa"},
                {"key": "7","value": "Soflusa"},
                {"key": "8","value": "Transportes Coletivos do Barreiro"},
                {"key": "9","value": "Vimeca Transportes"},
                {"key": "10","value": "Scotturb"},
                {"key": "11","value": "ID/JJ/HLM"},
                {"key": "12","value": "Isidoro Duarte"},
                {"key": "13","value": "Barraqueiro Transportes"},
                {"key": "14","value": "Joaquim Jerónimo"},
                {"key": "15","value": "Fertagus"},
                {"key": "16","value": "Metro Transportes do Sul"},
                {"key": "18","value": "Henrique Leonardo da Mota"},
                {"key": "21","value": "Cascais Próxima"},
                {"key": "23","value": "Portal VIVA"},
                {"key": "24","value": "Rodoviária do Tejo"},
                {"key": "34","value": "Câmara Municipal de Lisboa"},
                {"key": "41","value": "Viação Alvorada"},
                {"key": "42","value": "Rodoviária de Lisboa"},
                {"key": "43","value": "Transportes Sul do Tejo"},
                {"key": "44","value": "Alsa Todi"},
                {"key": "49","value": "Município de Oeiras"},
                {"key": "54","value": "Municipio de Setúbal"}
            ]
        },
        agency_url: {
            severity: "error",
        },
        agency_timezone: {
            severity: "error",
        },
        agency_lang: {
            severity: "ignore",
        },
        agency_phone: {
            severity: "error",
        },
        agency_fare_url: {
            severity: "error",
        },
        agency_email: {
            severity: "error",
        }
    },
    stops: {
        _file: "error",
        stop_id: {
            severity: "error",
        },
        stop_code: {
            severity: "error",
        },
        stop_name: {
            severity: "error",
        },
        stop_short_name: {
            severity: "ignore",
        },
        tts_stop_name: {
            severity: "ignore",
        },
        stop_desc: {
            severity: "ignore",
        },
        stop_lat: {
            severity: "error",
        },
        stop_lon: {
            severity: "error",
        },
        zone_id: {
            severity: "error",
        },
        stop_url: {
            severity: "ignore",
        },
        location_type: {
            severity: "error",
        },
        parent_station: {
            severity: "error",
        },
        stop_timezone: {
            severity: "ignore",
        },
        wheelchair_boarding: {
            severity: "error",
            options: ["0", "1", "2"]
        },
        level_id: {
            severity: "ignore",
        },
        platform_code: {
            severity: "error",
        },
        public_visible: {
            severity: "ignore",
            options: ["0", "1"]
        },
        has_stop_sign: {
            severity: "ignore",
            options: ["0", "1", "2", "3"]
        },
        has_shelter: {
            severity: "ignore",
            options: ["0", "1", "2", "3"]
        },
        shelter_code: {
            severity: "ignore",
        },
        shelter_maintainer: {
            severity: "ignore",
        },
        has_bench: {
            severity: "ignore",
            options: ["0", "1", "2", "3"]
        },
        has_network_map: {
            severity: "ignore",
            options: ["0", "1", "2", "3"]
        },
        has_schedules: {
            severity: "ignore",
            options: ["0", "1", "2", "3"]
        },
        has_pip_real_time: {
            severity: "ignore",
            options: ["0", "1", "2"]
        },
        has_tariffs_information: {
            severity: "ignore",
            options: ["0", "1", "3", "4"]
        },
        region_id: {
            severity: "error",
            options: ["PT170", "P185", "PT16B"]
        },
        municipality_id: {
            severity: "error",
            options: ["1502","1503","1504","1115","1105","1106","1107","1109","1506","1507","1116","1110","1508","1510","1111","1511","1512","1114","0712","1102","1112"]
        },
        parish_id: {
            severity: "error",
        }
    },
    routes: {
        _file: "error",
        line_id: {
            severity: "error",
        },
        line_short_name: {
            severity: "error",
        },
        line_long_name: {
            severity: "error",
        },
        route_id: {
            severity: "error",
        },
        agency_id: {
            severity: "error",
        },
        route_short_name: {
            severity: "error",
        },
        route_long_name: {
            severity: "error",
        },
        route_desc: {
            severity: "ignore",
        },
        route_remarks: {
            severity: "ignore",
        },
        route_type: {
            severity: "error",
            options: ["0", "1", "2", "3", "4", "5", "6", "7", "11", "12"]
        },
        path_type: {
            severity: "error",
            options: ["1", "2", "3"]
        },
        circular: {
            severity: "error",
            options: ["0", "1"]
        },
        school: {
            severity: "error",
            options: ["0", "1"]
        },
        route_url: {
            severity: "ignore",
        },
        route_color: {
            severity: "error",
        },
        route_text_color: {
            severity: "error",
        },
        continuous_pickup: {
            severity: "error",
            options: ["0", "1"]
        },
        continuous_drop_off: {
            severity: "error",
            options: ["0", "1"]
        }
    },
    trips: {
        _file: "error",
        route_id: {
            severity: "error",
        },
        pattern_id: {
            severity: "error",
        },
        service_id: {
            severity: "error",
        },
        trip_id: {
            severity: "error",
        },
        trip_headsign: {
            severity: "error",
        },
        trip_short_name: {
            severity: "ignore",
        },
        direction_id: {
            severity: "error",
        },
        block_id: {
            severity: "ignore",
        },
        shape_id: {
            severity: "error",
        },
        wheelchair_accessible: {
            severity: "error",
        },
        bikes_allowed: {
            severity: "error",
        },
        stop_sequence: {
            severity: "error",
        },
        direction_pattern_id_match: {
            severity: "error",
        },
        trip_id_limit_characters: {
            severity: "error",
        }
    },
    stop_times: {
        _file: "error",
        trip_id: {
            severity: "error",
        },
        arrival_time: {
            severity: "error",
        },
        departure_time: {
            severity: "error",
        },
        stop_id: {
            severity: "error",
        },
        stop_sequence: {
            severity: "error",
        },
        stop_headsign: {
            severity: "forbidden",
        },
        pickup_type: {
            severity: "error",
        },
        drop_off_type: {
            severity: "error",
        },
        continuous_pickup: {
            severity: "error",
        },
        continuous_drop_off: {
            severity: "error",
        },
        shape_dist_traveled: {
            severity: "error",
        },
        start_pickup_drop_off_window: {
            severity: "ignore",
        },
        end_pickup_drop_off_window: {
            severity: "ignore",
        },
        pickup_booking_rule_id: {
            severity: "ignore",
        },
        drop_off_booking_rule_id: {
            severity: "ignore",
        },
        timepoint: {
            severity: "error",
        },
        zone_1: {
            severity: "ignore",
        },
        zone_2: {
            severity: "ignore",
        },
        zone_3: {
            severity: "ignore",
        }
    },
    calendar: {
        _file: "error",
        service_id: {
            severity: "error",
        },
        monday: {
            severity: "error",
        },
        tuesday: {
            severity: "error",
        },
        wednesday: {
            severity: "error",
        },
        thursday: {
            severity: "error",
        },
        friday: {
            severity: "error",
        },
        saturday: {
            severity: "error",
        },
        sunday: {
            severity: "error",
        },
        start_date: {
            severity: "error",
        },
        end_date: {
            severity: "error",
        }
    },
    calendar_dates: {
        _file: "error",
        service_id: {
            severity: "error",
        },
        date: {
            severity: "error",
        },
        exception_type: {
            severity: "error",
        }
    },
    vehicles: {
        _file: "error",
        vehicle_id: {
            severity: "error",
        },
        agency_id: {
            severity: "error",
        },
        license_plate: {
            severity: "error",
        },
        make: {
            severity: "error",
        },
        model: {
            severity: "error",
        },
        owner: {
            severity: "error",
        },
        registration_date: {
            severity: "error",
        },
        available_seats: {
            severity: "error",
        },
        available_standing: {
            severity: "error",
        },
        typology: {
            severity: "error",
        },
        propulsion: {
            severity: "error",
            options: ["1", "2", "3", "4", "5", "6", "7", "8"]
        },
        emission: {
            severity: "error",
            options: ["1", "2", "3", "4", "5", "6"]
        },
        climatization: {
            severity: "error",
            options: ["0", "1"]
        },
        wheelchair: {
            severity: "error",
            options: ["0", "1"]
        },
        lowered_floor: {
            severity: "error",
            options: ["0", "1", "2"]
        },
        ramp: {
            severity: "error",
            options: ["0", "1", "2", "3"]
        },
        kneeling: {
            severity: "error",
            options: ["0", "1", "2"]
        },
        static_information: {
            severity: "error",
            options: ["0", "1"]
        },
        onboard_monitor: {
            severity: "error",
            options: ["0", "1"],
        },
        front_display: {
            severity: "error",
            options: ["0", "1"]
        },
        rear_display: {
            severity: "error",
            options: ["0", "1", "2"]
        },
        side_display: {
            severity: "error",
            options: ["0", "1", "2"]
        },
        internal_sound: {
            severity: "error",
            options: ["0", "1"]
        },
        external_sound: {
            severity: "error",
            options: ["0", "1"]
        },
        consumption_meter: {
            severity: "error",
            options: ["0", "1"]
        },
        bicycles: {
            severity: "error",
            options: ["0", "1"]
        },
        passenger_counting: {
            severity: "error",
            options: ["0", "1"]
        },
        video_surveillance: {
            severity: "error",
            options: ["0", "1"]
        }
    },
    fare_attributes: {
        _file: "error",
        fare_id: {
            severity: "error",
        },
        price: {
            severity: "error",
        },
        currency_type: {
            severity: "error",
        },
        payment_method: {
            severity: "error",
            options: ["0"]
        },
        transfers: {
            severity: "error",
            options: ["0"]
        },
        agency_id: {
            severity: "error",
        },
        transfer_duration: {
            severity: "error",
        }
    },
    fare_rules: {
        _file: "error",
        fare_id: {
            severity: "error",
        },
        route_id: {
            severity: "error",
        },
        origin_id: {
            severity: "forbidden",
        },
        destination_id: {
            severity: "forbidden",
        },
        contains_id: {
            severity: "forbidden",
        }
    },
    shapes: {
        _file: "error",
        shape_id: {
            severity: "error",
        },
        shape_pt_lat: {
            severity: "error",
        },
        shape_pt_lon: {
            severity: "error",
        },
        shape_pt_sequence: {
            severity: "error",
        },
        shape_dist_traveled: {
            severity: "error",
        }
    },
    frequencies: {
        _file: "forbidden",
        trip_id: {
            severity: "error",
        },
        start_time: {
            severity: "error",
        },
        end_time: {
            severity: "error",
        },
        headway_secs: {
            severity: "error",
        },
        exact_times: {
            severity: "error",
        }
    },
    transfers: {
        _file: "ignore",
        from_stop_id: {
            severity: "ignore",
        },
        to_stop_id: {
            severity: "ignore",
        },
        transfer_type: {
            severity: "ignore",
        },
        min_transfer_time: {
            severity: "ignore",
        }
    },
    pathways: {
        _file: "ignore",
        pathway_id: {
            severity: "ignore",
        },
        from_stop_id: {
            severity: "ignore",
        },
        to_stop_id: {
            severity: "ignore",
        },
        pathway_mode: {
            severity: "ignore",
        },
        is_bidirectional: {
            severity: "ignore",
        },
        length: {
            severity: "ignore",
        },
        traversal_time: {
            severity: "ignore",
        },
        stair_count: {
            severity: "ignore",
        },
        max_slope: {
            severity: "ignore",
        },
        min_width: {
            severity: "ignore",
        },
        signposted_as: {
            severity: "ignore",
        },
        reversed_signposted_as: {
            severity: "ignore",
        }
    },
    levels: {
        _file: "ignore",
        level_id: {
            severity: "ignore",
        },
        level_index: {
            severity: "ignore",
        },
        level_name: {
            severity: "ignore",
        }
    },
    feed_info: {
        _file: "error",
        feed_type: {
            severity: "error",
            options: ["0"]
        },
        feed_publisher_name: {
            severity: "error",
        },
        feed_publisher_url: {
            severity: "error",
        },
        feed_lang: {
            severity: "error",
            options: ["pt"]
        },
        default_lang: {
            severity: "error",
        },
        feed_start_date: {
            severity: "error",
        },
        feed_end_date: {
            severity: "error",
        },
        feed_version: {
            severity: "error",
        },
        feed_remarks: {
            severity: "ignore",
        },
        feed_contact_email: {
            severity: "error",
        },
        feed_contact_url: {
            severity: "error",
        }
    },
    translations: {
        _file: "ignore",
        table_name: {
            severity: "ignore",
        },
        field_name: {
            severity: "ignore",
        },
        language: {
            severity: "ignore",
        },
        translation: {
            severity: "ignore",
        },
        record_id: {
            severity: "ignore",
        },
        record_sub_id: {
            severity: "ignore",
        },
        field_value: {
            severity: "ignore",
        }
    },
    attributions: {
        _file: "ignore",
        attribution_id: {
            severity: "ignore",
        },
        agency_id: {
            severity: "ignore",
        },
        route_id: {
            severity: "ignore",
        },
        trip_id: {
            severity: "ignore",
        },
        organization_name: {
            severity: "ignore",
        },
        is_producer: {
            severity: "ignore",
        },
        is_operator: {
            severity: "ignore",
        },
        is_authority: {
            severity: "ignore",
        },
        attribution_url: {
            severity: "ignore",
        },
        attribution_email: {
            severity: "ignore",
        },
        attribution_phone: {
            severity: "ignore",
        }
    }
}