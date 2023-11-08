type OktaUser = {
  id: string,
  nrcOrganization: string,
  read: boolean,
  write: boolean,
};

type OldGroupConfig = {
  [environment: string]: {
    read: string,
    write: string,
  } 
};

type NewGroupConfig = {
  [environment: string]: {
    [nrcOrganization: string]: {
      read: string[],
      write: string[],
    },
  },
};

const env = 'uat';
const dryRun = true;
const oktaDomain = 'nrc.okta.com';
const oktaApiKey = 'SSWS xxx';

// Map of environment to app id
const appIdMap: Record<string, string> = {
  uat: '0oa1md8zcb820pDUP1d8',
  prod: '0oa1mfoz6adtVAZv01d8',
};

// Map of group name to group id
// This is extended as groups are created
const groupIdMap: Record<string, string> = {
  'App_NRC_Core_UAT_Read_Group': '00g1mnexikp9SW2ut1d8',
  'App_NRC_Core_UAT_Write_Group': '00g1mnexd6e1Q2oh41d8',
  'App_NRC_Core_Read_Group': '00g1mnexksv7gJwD81d8',
  'App_NRC_Core_Write_Group': '00g1mnetclnuL6J7J1d8',
};

// Map of environment to permission to old groups
const oldGroups: OldGroupConfig = {
  'uat': {
    read: 'App_NRC_Core_UAT_Read_Group',
    write: 'App_NRC_Core_UAT_Write_Group',
  },
  'prod': {
    read: 'App_NRC_Core_Read_Group',
    write: 'App_NRC_Core_Write_Group',
  },
};

// Map of environment to NRC Org to permission to new groups
const nrcOrgToGroupMap: NewGroupConfig = {
  'uat': {
    'NRC Bangladesh': {
      read: ['App__NRC_CORE__UAT__Bangladesh__READ'],
      write: ['App__NRC_CORE__UAT__Bangladesh__WRITE'],
    },
    'NRC Burkina Faso': {
      read: ['App__NRC_CORE__UAT__Burkina_Faso__READ'],
      write: ['App__NRC_CORE__UAT__Burkina_Faso__WRITE'],
    },
    'NRC Central West Africa': {
      read: ['App__NRC_CORE__UAT__Burkina_Faso__READ'],
      write: ['App__NRC_CORE__UAT__Burkina_Faso__WRITE'],
    },
    'NRC Cameroon': {
      read: ['App__NRC_CORE__UAT__Cameroon__READ'],
      write: ['App__NRC_CORE__UAT__Cameroon__WRITE'],
    },
    'NRC Central West Africa Regional Office': {
      read: ['App__NRC_CORE__UAT__Cameroon__READ'],
      write: ['App__NRC_CORE__UAT__Cameroon__WRITE'],
    },
    'NRC Central African Republic': {
      read: ['App__NRC_CORE__UAT__Central_African_Republic__READ'],
      write: ['App__NRC_CORE__UAT__Central_African_Republic__WRITE'],
    },
    'NRC Colombia': {
      read: ['App__NRC_CORE__UAT__Colombia__READ'],
      write: ['App__NRC_CORE__UAT__Colombia__WRITE'],
    },
    'NRC DR Congo': {
      read: ['App__NRC_CORE__UAT__DR_Congo__READ'],
      write: ['App__NRC_CORE__UAT__DR_Congo__WRITE'],
    },
    'NRC Iran': {
      read: ['App__NRC_CORE__UAT__Iran__READ'],
      write: ['App__NRC_CORE__UAT__Iran__WRITE'],
    },
    'NRC Kenya': {
      read: ['App__NRC_CORE__UAT__Kenya__READ'],
      write: ['App__NRC_CORE__UAT__Kenya__WRITE'],
    },
    'NRC Libya': {
      read: ['App__NRC_CORE__UAT__Libya__READ'],
      write: ['App__NRC_CORE__UAT__Libya__WRITE'],
    },
    'NRC Mozambique': {
      read: ['App__NRC_CORE__UAT__Mozambique__READ'],
      write: ['App__NRC_CORE__UAT__Mozambique__WRITE'],
    },
    'NRC North Central America': {
      read: ['App__NRC_CORE__UAT__North_Central_America__READ'],
      write: ['App__NRC_CORE__UAT__North_Central_America__WRITE'],
    },
    'NRC Head Office': {
      read: ['App__NRC_CORE__UAT__Test_Country__READ'],
      write: ['App__NRC_CORE__UAT__Test_Country__WRITE'],
    },
    'NRC Germany': {
      read: ['App__NRC_CORE__UAT__Test_Country__READ'],
      write: ['App__NRC_CORE__UAT__Test_Country__WRITE'],
    },
    'NRC Central and Eastern Europe Regional Office': {
      read: ['App__NRC_CORE__UAT__Test_Country__READ'],
      write: ['App__NRC_CORE__UAT__Test_Country__WRITE'],
    },
    'NRC Palestine': {
      read: ['App__NRC_CORE__UAT__Palestine__READ'],
      write: ['App__NRC_CORE__UAT__Palestine__WRITE'],
    },
    'NRC Middle East Regional Office': {
      read: ['App__NRC_CORE__UAT__Palestine__READ'],
      write: ['App__NRC_CORE__UAT__Palestine__WRITE'],
    },
    'NRC Somalia': {
      read: ['App__NRC_CORE__UAT__Somalia__READ'],
      write: ['App__NRC_CORE__UAT__Somalia__WRITE'],
    },
    'NRC South Sudan': {
      read: ['App__NRC_CORE__UAT__South_Sudan__READ'],
      write: ['App__NRC_CORE__UAT__South_Sudan__WRITE'],
    },
    'NRC Sudan': {
      read: ['App__NRC_CORE__UAT__Sudan__READ'],
      write: ['App__NRC_CORE__UAT__Sudan__WRITE'],
    },
    'NRC EAY Regional Office': {
      read: ['App__NRC_CORE__UAT__Sudan__READ'],
      write: ['App__NRC_CORE__UAT__Sudan__WRITE'],
    },
    'NRC East Africa and Yemen Regional Office': {
      read: ['App__NRC_CORE__UAT__Sudan__READ'],
      write: ['App__NRC_CORE__UAT__Sudan__WRITE'],
    },
    'NRC Syria': {
      read: ['App__NRC_CORE__UAT__Syria__READ'],
      write: ['App__NRC_CORE__UAT__Syria__WRITE'],
    },
    'NRC Tanzania': {
      read: ['App__NRC_CORE__UAT__Tanzania__READ'],
      write: ['App__NRC_CORE__UAT__Tanzania__WRITE'],
    },
    'NRC Uganda': {
      read: ['App__NRC_CORE__UAT__Uganda__READ'],
      write: ['App__NRC_CORE__UAT__Uganda__WRITE'],
    },
    'NRC Mali': {
      read: ['App__NRC_CORE__UAT__Mali__READ'],
      write: ['App__NRC_CORE__UAT__Mali__WRITE'],
    },
  },
  'prod': {
    'NRC Bangladesh': {
      read: ['App__NRC_CORE__PROD__Bangladesh__READ'],
      write: ['App__NRC_CORE__PROD__Bangladesh__WRITE'],
    },
    'NRC Cameroon': {
      read: ['App__NRC_CORE__PROD__Cameroon__READ'],
      write: ['App__NRC_CORE__PROD__Cameroon__WRITE'],
    },
    'NRC Head Office': {
      read: ['App__NRC_CORE__PROD__Test_Country__READ', 'App__NRC_CORE__PROD__Uganda__READ'],
      write: ['App__NRC_CORE__PROD__Test_Country__WRITE', 'App__NRC_CORE__PROD__Uganda__WRITE'],
    },
    'NRC Palestine': {
      read: ['App__NRC_CORE__PROD__Palestine__READ'],
      write: ['App__NRC_CORE__PROD__Palestine__WRITE'],
    },
    'NRC Uganda': {
      read: ['App__NRC_CORE__PROD__Uganda__READ'],
      write: ['App__NRC_CORE__PROD__Uganda__WRITE'],
    },
    'NRC South Sudan': {
      read: ['App__NRC_CORE__PROD__Uganda__READ'],
      write: ['App__NRC_CORE__PROD__Uganda__WRITE'],
    },
  },
};

// From a group name, get or create the group, and add it to the group id map
const createGroup = async (name: string): Promise<void> => {
  console.info(`Creating group ${name}`);

  if (dryRun) {
    return;
  }

  const getResp = await fetch(
    `https://${oktaDomain}/api/v1/groups?q=${name}`,
    {
      method: 'GET',
      headers: {
        Authorization: oktaApiKey, 
        'Content-Type': 'application/json',
      },
    }
  );

  if (getResp.status !== 200) {
    throw new Error(`Failed to get group ${name}`);
  }
  const getData = await getResp.text();
  const groups = JSON.parse(getData);
  if (groups.length > 0) {
    console.info(`Group ${name} already exists`);
    groupIdMap[name] = groups[0].id;
    return;
  }

  const resp = await fetch(
    `https://${oktaDomain}/api/v1/groups`,
    {
      method: 'POST',
      headers: {
        Authorization: oktaApiKey, 
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        profile: {
          name,
        },
        type: 'APP_GROUP',
      }),
    }
  );

  if (resp.status !== 200) {
    throw new Error(`Failed to create group ${name}`);
  }
  const data = await resp.text();
  const group = JSON.parse(data);
  groupIdMap[name] = group.id;

  // assign to app
  const assignResp = await fetch(
    `https://${oktaDomain}/api/v1/apps/${appIdMap[env]}/groups/${group.id}`,
    {
      method: 'PUT',
      headers: {
        Authorization: oktaApiKey, 
      },
    },
  );

  if (assignResp.status !== 200) {
    throw new Error(`Failed to assign group ${name} to app`);
  }
};

// For a list of group names, get all users in those groups
const getUsersInGroups = async (groups: OldGroupConfig[string]): Promise<OktaUser[]> => {

  const getUsersInGroup = async (groupId: string, perm: 'read' | 'write'): Promise<OktaUser[]> => {
    const resp = await fetch(
      `https://${oktaDomain}/api/v1/groups/${groupId}/users`,
      {
        method: 'GET',
        headers: {
          Authorization: oktaApiKey, 
        }
      }
    );

    if (resp.status !== 200) {
      console.log(resp.status)
      throw new Error(`Failed to get users in group ${groupId}`);
    }
    
    const data = await resp.text();
    const users = JSON.parse(data);
    return users.map((user: any) => ({
      id: user.id,
      nrcOrganization: user['profile']['nrcOrganisation'],
      read: perm === 'read',
      write: perm === 'write',
    }));
  };

  const users: Record<string, OktaUser> = {}; 

  for (const [perm, group] of Object.entries(groups)) {
    const u = await getUsersInGroup(groupIdMap[group], perm as 'read' | 'write');
    for (const user of u) {
      if (users[user.id]) {
        users[user.id].read = users[user.id].read || user.read;
        users[user.id].write = users[user.id].write || user.write;
      } else {
        users[user.id] = user;
      }
    }
  }

  return Object.values(users);
};

// Add a user to a group using the group name
const addUserToGroup = async (user: OktaUser, group: string): Promise<void> => {
  console.info(`Adding user ${user.id} (NRC Org: ${user.nrcOrganization}) to group ${group}`);

  if (dryRun) {
    return;
  }

  const resp = await fetch(
    `https://${oktaDomain}/api/v1/groups/${groupIdMap[group]}/users/${user.id}`,
    {
      method: 'PUT',
      headers: {
        Authorization: oktaApiKey, 
      },
    },
  );

  if (resp.status !== 204) {
    throw new Error(`Failed to add user ${user.id} to group ${group}`);
  }
};

const migrate = async (environment: 'uat' | 'prod') => {
  // Get environment specific groups
  const nrcOrgToGroupMapEnv = nrcOrgToGroupMap[environment];
  const oldGroupsEnv = oldGroups[environment];

  // Get new groups as a list
  const newGroups = Object.values(nrcOrgToGroupMapEnv).reduce<string[]>((acc, groups) => {
    acc.push(...groups.read);
    acc.push(...groups.write);
    return acc;
  }, []);

  // Create new groups
  for (const group of newGroups) {
    await createGroup(group);
  }

  // Get users in old groups
  const users = await getUsersInGroups(oldGroupsEnv);

  // console.log(JSON.stringify(users, null, 2))

  // Iterate over users
  for (const user of users) {
    // Get new groups for user
    const groups = nrcOrgToGroupMap[environment][user.nrcOrganization];

    if (!groups) {
      console.warn(`User ${user.id} (NRC Org: ${user.nrcOrganization}) not configured in environment ${environment}`);
      // throw new Error(`User ${user.id} (NRC Org: ${user.nrcOrganization}) has no groups in environment ${environment}`);
      continue;
    }

    // If user had read or write access, add them to read groups
    if (user.read || user.write) {
      for (const group of groups.read) {
        await addUserToGroup(user, group);
      }
    }

    // If user had write access, add them to write groups
    if (user.write) {
      for (const group of groups.write) {
        await addUserToGroup(user, group);
      }
    }
  }
};

migrate(env);