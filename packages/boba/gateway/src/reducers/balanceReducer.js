/*
Copyright 2019-present OmiseGO Pte Ltd

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License. */

const initialState = {
  layer1: [],
  layer2: [],
  l1LpBalanceWeiString:'',
  l2LpBalanceWeiString:'',
  l1FeeRate: 0,
  l2FeeRate: 0,
  l1GasFee: 0,
  l2FeeBalance: 0,
  l1lpLiquidity: 0,
  l2lpLiquidity: 0
}

function balanceReducer(state = initialState, action) {
  switch (action.type) {
    case 'BALANCE/GET/SUCCESS':
      const { layer1, layer2 } = action.payload
      return { ...state, layer1, layer2 }
    case 'FETCH/L1LP/BALANCE/SUCCESS':
      return {
        ...state, 
        l1LpBalanceWeiString: action.payload
      }
    case 'FETCH/L2LP/BALANCE/SUCCESS':
      return {
        ...state, 
        l2LpBalanceWeiString: action.payload
      }
    case 'FETCH/L1TOTALFEERATE/SUCCESS':
      return {
        ...state, 
        l1FeeRate: action.payload
      }
    case 'FETCH/L2TOTALFEERATE/SUCCESS':
      return {
        ...state, 
        l2FeeRate: action.payload
      }
    case 'FETCH/FASTEXIT/COST/SUCCESS':
      return {
        ...state, 
        l1GasFee: action.payload
      }
    case 'FETCH/L2FEE/BALANCE/SUCCESS':
      return {
        ...state, 
        l2FeeBalance: action.payload
      }
    case 'FETCH/L1LP/LIQUIDITY/SUCCESS':
      return {
        ...state, 
        l1lpLiquidity: action.payload
      }
    case 'FETCH/L2LP/LIQUIDITY/SUCCESS':
      return {
        ...state, 
        l2lpLiquidity: action.payload
      }
    case 'BALANCE/L1/RESET':
      return {
        ...state, 
        l1LpBalanceWeiString: '',
        l1FeeRate: 0,
        l1GasFee: 0,
        l2FeeBalance: 0,
        l1lpLiquidity: 0,
      }
    case 'BALANCE/L2/RESET':
      return {
        ...state, 
        l2LpBalanceWeiString: '',
        l2FeeRate: 0,
        l2lpLiquidity: 0,
      }
    default:
      return state
  }
}

export default balanceReducer
