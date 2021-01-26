interface AccountStateProps {
    User: object | null;
    Token: string | null;
}

const defaultState: AccountStateProps = {
    User: null,
    Token: null,
}

const Accounts = {
    state: defaultState,
    mutations: {},
    actions: {},
};

export default Accounts;