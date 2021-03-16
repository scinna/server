import React                     from 'react';
import {Token}                   from "../../types/Token";
import {Token as TokenComponent} from './Token';
import {useApiCall}              from "../../utils/useApi";
import {Loader}                  from "../Loader";

export function TokenLister() {
    const tokens = useApiCall<Token[]>({ url: '/api/account/tokens' });

    return <div className="tokenLister">
        {
            tokens.status === 'pending'
            &&
                <Loader/>
        }
        {
            tokens.status === 'success'
            &&
                tokens.data.map(token => <TokenComponent key={token.Token} token={token}/>)
        }
    </div>;
}