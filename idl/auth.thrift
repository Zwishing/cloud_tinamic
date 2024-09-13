namespace go base.auth

include "base.thrift"

enum RoleCode{
    SUPERADMIN = 1,
    ADMIN = 2,
    USER = 3,
    GUEST = 4,
}

struct Role {
     1: i64 id,
     2: i64 parent_id,
     3: string name,
     4: RoleCode role_code,
}

struct AuthResquest{
    1: string sub;
    2: string obj;
    3: string act
}

struct AuthResponse{
    1: base.BaseResp base;
    2: bool allow;
}
struct EditResponse{
    1: base.BaseResp base;
    2: bool success;
}

service AuthService{
    AuthResponse Auth(1:AuthResquest req),
    EditResponse AddPolicy(1:AuthResquest req)
    EditResponse RemovePolicy(1:AuthResquest req)
    EditResponse UpdatePolicy(1:AuthResquest req)
}




