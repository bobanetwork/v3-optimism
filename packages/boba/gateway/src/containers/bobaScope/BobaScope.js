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

import React, { useState } from 'react'
import { batch, useDispatch } from 'react-redux'
import { isEqual, orderBy } from 'lodash'
import { useSelector } from 'react-redux'
import DatePicker from 'react-datepicker'
import "react-datepicker/dist/react-datepicker.css"

import { useMediaQuery, useTheme } from '@material-ui/core'

import { setActiveDataTab } from 'actions/uiAction'
import { fetchSevens } from 'actions/networkAction'

import { selectActiveDataTab } from 'selectors/uiSelector'
import { selectSevens } from 'selectors/dataSelector'

import Tabs from 'components/tabs/Tabs'
import Input from 'components/input/Input'

import Sevens from './Sevens'

import * as styles from './Transactions.module.scss'
import * as S from './History.styles'

import useInterval from 'util/useInterval'
import PageHeader from 'components/pageHeader/PageHeader'

import { POLL_INTERVAL } from 'util/constant'

function BobaScope() {

  const theme = useTheme()
  const isMobile = useMediaQuery(theme.breakpoints.down('md'));

  const dispatch = useDispatch()

  const now = new Date()
  const last_week = new Date(now.getFullYear(), now.getMonth(), now.getDate() - 7)

  const [startDate, setStartDate] = useState(last_week)
  const [endDate, setEndDate] = useState(now)

  const [searchData, setSearchData] = useState('')

  const activeTab = useSelector(selectActiveDataTab, isEqual)

  const unorderedSevens = useSelector(selectSevens, isEqual)

  //sort transactions by timeStamp
  const orderedSevens = orderBy(unorderedSevens, i => i.timeStamp, 'desc')

  const sevens = orderedSevens

  useInterval(() => {
    batch(()=>{dispatch(fetchSevens())})
  }, POLL_INTERVAL)

  return (
    <>
      <PageHeader title="BobaScope" />

      <S.Header>
        <div className={styles.searchInput}>
          <Input
            size='small'
            placeholder='Search by hash'
            value={searchData}
            onChange={i=>{setSearchData(i.target.value)}}
            className={styles.searchBar}
          />
        </div>
      </S.Header>
      <div className={styles.data}>
        <div className={styles.section}>
          <Tabs
            onClick={tab => {dispatch(setActiveDataTab(tab))}}
            activeTab={activeTab}
            tabs={['Seven Day Queue']}
          />

          {activeTab === 'Seven Day Queue' && (
            <Sevens
              searchData={searchData}
              sevens={sevens}
            />
          )}

        </div>
      </div>
    </>
  );
}

export default React.memo(BobaScope)
