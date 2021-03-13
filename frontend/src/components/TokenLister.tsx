import React                     from 'react';
import {useState}                from "react";
import {Token}                   from "../types/Token";
import {Token as TokenComponent} from './Token';
import useAsyncEffect            from "use-async-effect";
import {useToken}                from "../utils/TokenProvider";

export function TokenLister() {
    const { token } = useToken();
    const [tokens, setTokens] = useState<Token[]>([]);

    useAsyncEffect(async () => {
        const response = await fetch('/api/account/tokens', {
            headers: { Authorization: 'Bearer ' + token }
        })

        if (!response.ok) {
            // @TODO something
            console.log("nope! ", response.status)
            return
        }

        const data = await response.json()
        setTokens(data);
    }, [])

    return <div className="tokenLister">
        {
            tokens.map(token => <TokenComponent key={token.Token} token={token}/>)
        }
    </div>;
}