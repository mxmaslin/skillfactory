-- ФП1.3.1

select
    count(distinct u.id)
from 
    case10.users u
join
    case10.addresses a 
    on u.id = a.addressable_id 
join 
    case10.cities c 
    on c.id = a.city_id 
join
    case10.regions r 
    on r.id = c.region_id 
join 
    case10.countries ct
    on ct.id = r.country_id 
where 
    ct.name = 'Russia' 

-- ФП1.3.2

select
    count(distinct id)
from
    case10.users u 
where
    u.phone::text !~* '^(73|74|78|79)'

-- ФП1.3.4

with country as (
    select
        u.id id
    from 
        case10.users u
    join
        case10.addresses a
        on u.id = a.addressable_id
    join
        case10.cities c
        on a.city_id = c.id
    join
        case10.regions r
        on c.region_id = r.id
    join
        case10.countries co
        on r.country_id = co.id
    where co.name = 'Russia'
),
phone as (
    select
        u.id id
    from
        case10.users u 
    where
        u.phone::text !~* '^(73|74|78|79)'
),
ip as (
    select
        u.id id,
        substring(u.last_sign_in_ip,'\d+')::int*(256^3) + 
        substring(u.last_sign_in_ip ,'\d+.(\d+)')::int*(256^2) + 
        substring(u.last_sign_in_ip ,'\d+.\d+.(\d+)')::int*(256^1) + 
        substring(u.last_sign_in_ip ,'\d+.\d+.\d+.(\d+)')::int*(256^0) as last_sign_in_ip_key
    from
        case10.users u
),
ip_address as (
    select 
        ip.id id
    from
        ip
    join
        case10.ip2location_db1 ip2loc
        on ip.last_sign_in_ip_key between ip2loc.ip_from and ip2loc.ip_to
    where
        ip2loc.country_name = 'Russian Federation'
),
from_russia as (
    select id from country
    union
    select id from phone
    union
    select id from ip_address
)
select count(distinct id) from from_russia

-- ФП1.3.5

with ip as
(
    select
        *,
        last_sign_in_ip,
        substring(last_sign_in_ip,'\d+')::int*(256^3)+
        substring(last_sign_in_ip,'\d+.(\d+)')::int*(256^2)+
        substring(last_sign_in_ip,'\d+.\d+.(\d+)')::int*(256^1)+
        substring(last_sign_in_ip,'\d+.\d+.\d+.(\d+)')::int*(256^0) as last_sign_in_ip_key
    from case10.users
),
user_address as
(
    select
        con.name, us.id
    from
        case10.countries con
    join
        case10.regions reg 
        on con.id = reg.country_id
    join
        case10.cities ci
        on ci.region_id = reg.id
    join case10.addresses adr
        on adr.city_id = ci.id
    join case10.users us
        on us.id = adr.addressable_id
    where con.name ='Russia' 
),
user_phone as
(
    select
        Id, phone, 'Russian Federation' Country_name
    from
        case10.users
    where
        phone::text ~* '^(73|74|78|79)'
),
russian_users_ip_joined as
(
    select 
        ip.id, ip2.country_name
    from
        ip 
    join
        case10.ip2location_db1 ip2 
        on ip.last_sign_in_ip_key >= ip2.ip_from and ip.last_sign_in_ip_key <= ip2.ip_to
    where
        ip2.Country_name ='Russian Federation'
),
from_russia as
(
    select
        id
    from
        user_address
    union
    select
        id
    from
        user_phone
    union
    select
        id
    from
        russian_users_ip_joined
), 
user_regist as 
(
    select
        id, date_trunc('month', created_at)::date dd
    from
        case10.users 
    where
        id in (select * from from_russia)
),
id_by_purshase as
(
    select
        id, user_id, date_trunc('month',purchased_at)::date dd, state
    from
        case10.carts
    where
        user_id in (select * from from_russia) and state = 'successful'
),
cart_by_months as
(
    select
        user_regist.dd, count(id_by_purshase.id) Amount_user, count(user_regist.id) carts_amount
    from
        user_regist
    join
        id_by_purshase
        on user_regist.dd = id_by_purshase.dd and user_regist.id = id_by_purshase.user_id
    group by
        user_regist.dd
    order by 1 desc
),
amount_user_table as 
(
    select
        count (id) users_amount, date_trunc('month',created_at)::date dd
    from
        case10.users 
    where
        id in (select * from from_russia) 
    group by
        dd
    having
        count(id) >= 100
)
select 
    aut.dd, cbm.carts_amount, aut.users_amount, cbm.carts_amount * 1.0 / aut.users_amount as conversion
from
    amount_user_table aut
join
    cart_by_months cbm
    on aut.dd = cbm.dd
order by 4 desc
