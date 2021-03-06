create table if not exists univercity
(
    first_name                      text    not null,
    second_name                     text    not null,
    third_name                      text    not null,
    latin_first_name                text    not null,
    latin_second_name               text    not null,
    gender                          text not null,
    date_of_birth                   integer not null,
    is_foreigner                    boolean not null,
    nationality                     text    not null,
    family_status                   text    not null,
    region                          text    not null,
    post_index                      text    not null,
    birth_place                     text    not null,
    register_address                text    not null,
    home_address                    text    not null,
    live_address                    text    not null,
    phone_number                    text    not null,
    home_phone_number               text    not null,
    document_type                   text    not null,
    serial_and_passport_number      text not null,
    issued_by                       text    not null,
    when_issud                      integer not null,
    valid_until                     integer not null,
    personal_nomer                  text    not null,
    form_studing                    text not null,
    faculty                         text    not null,
    specialization                  text    not null,
    type_education                  text    not null,
    educational_institution         text    not null,
    education_specialization        text    not null,
    graduated_institution_year      integer not null,
    subject_1_name                  text    not null,
    subject_2_name                  text    not null,
    subject_3_name                  text    not null,
    subject_1_rating                real    not null,
    subject_2_rating                real    not null,
    subject_3_rating                real    not null,
    subject_1_setificate            text    not null,
    subject_2_setificate            text    not null,
    subject_3_setificate            text    not null,
    forigen_lang_name               text    not null,
    forigen_lang_rating             real    not null,
    document_type_education         text    not null,
    is_selskaya_shkola              boolean not null,
    is_gorodskaya                   boolean not null,
    a2019                           boolean not null,
    medal                           boolean not null,
    ssuz                            boolean not null,
    ptu                             boolean not null,
    document_education_number_nomer text    not null,
    document_education_at           integer not null,
    average_mark                    real    not null,
    olimp                           boolean not null,
    spec_fond_rb                    boolean not null,
    olimp_sport                     boolean not null,
    perviy_razrad                   boolean not null,
    lgotniy_kredit                  boolean not null,
    hostel                          boolean not null,
    disabled                        boolean not null,
    chernobil                       boolean not null,
    additional_information          text    not null,
    email                           text not null,
    privileges                      boolean not null,
    info_privileges                 text    not null,
    has_father                      boolean not null,
    second_name_father              text    not null,
    third_name_father               text    not null,
    adress_father                   text    not null,
    work_adress_father              text    not null,
    contact_father                  text    not null,
    has_mother                      boolean not null,
    second_mother_name              text    not null,
    mother_first_name               text    not null,
    third_mother_name               text    not null,
    work_mother_address             text    not null,
    mother_contact_info             text    not null,
    owner_passport                  text    not null,
    owner_passport_number           text    not null,
    owner_passport_id               text    not null,
    owner_passport_kto_vidal        text    not null,
    owner_passport_kogda_vidal      integer not null,
    owner_passport_srok             integer not null,
    pasport_photo                   text    not null,
    ct_1_photo                      text    not null,
    ct_2_photo                      text    not null,
    ct_3_photo                      text    not null,
    ct_4_photo                      text    not null,
    owner_pasport_photo             text    not null,
    education_document_photo        text    not null,
    first_name_father               text not null,
    adress_mother                   text    not null,
    sum_marks                       real    not null
);

