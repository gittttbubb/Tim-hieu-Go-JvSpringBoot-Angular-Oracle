-- ==========================================
-- ROLES
-- ==========================================

INSERT INTO roles (
    id,
    name,
    display_name,
    description
)
VALUES (
    '11111111-1111-1111-1111-111111111111',
    'ADMIN',
    'Administrator',
    'System administrator'
);

INSERT INTO roles (
    id,
    name,
    display_name,
    description
)
VALUES (
    '22222222-2222-2222-2222-222222222222',
    'MANAGER',
    'Manager',
    'Manager role'
);

INSERT INTO roles (
    id,
    name,
    display_name,
    description
)
VALUES (
    '33333333-3333-3333-3333-333333333333',
    'USER',
    'User',
    'Standard user'
);

-- ==========================================
-- PERMISSIONS
-- ==========================================

INSERT INTO permissions VALUES (
    'p001',
    'USER',
    'USER_VIEW',
    'VIEW',
    'View users'
);

INSERT INTO permissions VALUES (
    'p002',
    'USER',
    'USER_CREATE',
    'CREATE',
    'Create user'
);

INSERT INTO permissions VALUES (
    'p003',
    'USER',
    'USER_UPDATE',
    'UPDATE',
    'Update user'
);

INSERT INTO permissions VALUES (
    'p004',
    'USER',
    'USER_DELETE',
    'DELETE',
    'Delete user'
);

INSERT INTO permissions VALUES (
    'p005',
    'USER',
    'USER_PERMISSION_VIEW',
    'VIEW',
    'View user permission overrides'
);

INSERT INTO permissions VALUES (
    'p006',
    'USER',
    'USER_PERMISSION_ASSIGN',
    'ASSIGN',
    'Assign user permission override'
);

INSERT INTO permissions VALUES (
    'p007',
    'USER',
    'USER_PERMISSION_REMOVE',
    'REMOVE',
    'Remove user permission override'
);

INSERT INTO permissions VALUES (
    'p008',
    'ROLE',
    'ROLE_VIEW',
    'VIEW',
    'View roles'
);

INSERT INTO permissions VALUES (
    'p009',
    'ROLE',
    'ROLE_CREATE',
    'CREATE',
    'Create role'
);

INSERT INTO permissions VALUES (
    'p010',
    'ROLE',
    'ROLE_UPDATE',
    'UPDATE',
    'Update role'
);

INSERT INTO permissions VALUES (
    'p011',
    'ROLE',
    'ROLE_DELETE',
    'DELETE',
    'Delete role'
);

INSERT INTO permissions VALUES (
    'p012',
    'ROLE',
    'ROLE_PERMISSION_VIEW',
    'VIEW',
    'View role permissions'
);

INSERT INTO permissions VALUES (
    'p013',
    'ROLE',
    'ROLE_PERMISSION_ASSIGN',
    'ASSIGN',
    'Assign permission to role'
);

INSERT INTO permissions VALUES (
    'p014',
    'ROLE',
    'ROLE_PERMISSION_REMOVE',
    'REMOVE',
    'Remove permission from role'
);

INSERT INTO permissions VALUES (
    'p015',
    'PERMISSION',
    'PERMISSION_VIEW',
    'VIEW',
    'View permissions'
);

INSERT INTO permissions VALUES (
    'p016',
    'AUDIT',
    'AUDIT_VIEW',
    'VIEW',
    'View audit logs'
);

-- ==========================================
-- ADMIN ROLE PERMISSIONS (ALL)
-- ==========================================

INSERT INTO role_permissions
SELECT
    '11111111-1111-1111-1111-111111111111',
    id,
    1,
    'ALL'
FROM permissions;

-- ==========================================
-- MANAGER ROLE PERMISSIONS (TEAM)
-- ==========================================

INSERT INTO role_permissions VALUES (
    '22222222-2222-2222-2222-222222222222',
    'p001',
    1,
    'TEAM'
);

INSERT INTO role_permissions VALUES (
    '22222222-2222-2222-2222-222222222222',
    'p002',
    1,
    'TEAM'
);

INSERT INTO role_permissions VALUES (
    '22222222-2222-2222-2222-222222222222',
    'p003',
    1,
    'TEAM'
);

INSERT INTO role_permissions VALUES (
    '22222222-2222-2222-2222-222222222222',
    'p005',
    1,
    'TEAM'
);

INSERT INTO role_permissions VALUES (
    '22222222-2222-2222-2222-222222222222',
    'p008',
    1,
    'TEAM'
);

INSERT INTO role_permissions VALUES (
    '22222222-2222-2222-2222-222222222222',
    'p015',
    1,
    'TEAM'
);

-- ==========================================
-- USER ROLE PERMISSIONS (OWN)
-- ==========================================

INSERT INTO role_permissions VALUES (
    '33333333-3333-3333-3333-333333333333',
    'p001',
    1,
    'OWN'
);

-- ==========================================
-- ADMIN USER
-- username : admin
-- password : admin123
-- ==========================================

INSERT INTO users (
    id,
    tenant_id,
    full_name,
    username,
    email,
    phone,
    password_hash,
    role_id,
    status,
    must_change_password,
    created_at,
    updated_at,
    created_by,
    password_changed_at
)
VALUES (
    'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
    '00000000-0000-0000-0000-000000000001',
    'System Administrator',
    'admin',
    'admin@example.com',
    '0900000000',
    '$2a$10$kwVEGmmqzRabkD99XwgMwuz7sERaBivMK1uemIWA1RIpOd1Utgh3y',
    '11111111-1111-1111-1111-111111111111',
    'ACTIVE',
    1,
    SYSTIMESTAMP,
    SYSTIMESTAMP,
    NULL,
    SYSTIMESTAMP
);

COMMIT;