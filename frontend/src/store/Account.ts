interface AccountStateProps {
    User: object | null;
    Token: string | null;
}

const defaultState: AccountStateProps = {
    User: null,
    Token: null,
}

const accounts = {
    state: defaultState,
    mutations: {},
    actions: {},
};

export default accounts;
