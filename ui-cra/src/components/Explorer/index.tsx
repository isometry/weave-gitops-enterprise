import { Box, IconButton } from '@material-ui/core';
import {
  Flex,
  Icon,
  IconType,
  RouterTab,
  SubRouterTabs,
} from '@weaveworks/weave-gitops';
// @ts-ignore
import styled from 'styled-components';
import { useQueryService } from '../../hooks/query';
import { Routes } from '../../utils/nav';
import { ContentWrapper } from '../Layout/ContentWrapper';
import { PageTemplate } from '../Layout/PageTemplate';
import AccessRulesDebugger from './AccessRulesDebugger';
import ExplorerTable from './ExplorerTable';
import {
  columnHeaderHandler,
  filterChangeHandler,
  useQueryState,
} from './hook';
import QueryBuilder from './QueryBuilder';

type Props = {
  className?: string;
};

// ?clusterName=management&name=flux-system&namespace=flux-system

function Explorer({ className }: Props) {
  const [queryState, setQueryState] = useQueryState({
    enableURLState: true,
    filters: [
      { label: 'Kustomizations', value: 'kind:Kustomization' },
      { label: 'Helm Releases', value: 'kind:HelmRelease' },
      {
        label: 'Failed',
        value: 'status:Failed',
      },
    ],
  });

  const { data, error, isFetching } = useQueryService({
    query: queryState.pinnedTerms.join(','),
    limit: queryState.limit,
    offset: queryState.offset,
    orderBy: `${queryState.orderBy} ${
      queryState.orderDescending ? 'desc' : 'asc'
    }`,
  });

  const handlePageForward = () => {
    setQueryState({
      ...queryState,
      offset: queryState.offset + queryState.limit,
    });
  };

  const handlePageBack = () => {
    setQueryState({
      ...queryState,
      offset: Math.max(0, queryState.offset - queryState.limit),
    });
  };

  return (
    <PageTemplate documentTitle="Explorer" path={[{ label: 'Explorer' }]}>
      <ContentWrapper
        errors={
          error
            ? // Hack to get the message to format correctly.
              // The ContentWrapper API should be simplified to support things other than ListError.
              [{ clusterName: 'Error', message: error?.message }]
            : undefined
        }
      >
        <div className={className}>
          <SubRouterTabs rootPath={`${Routes.Explorer}/query`}>
            <RouterTab name="Query" path={`${Routes.Explorer}/query`}>
              <>
                <Flex align>
                  <QueryBuilder
                    busy={isFetching}
                    disabled={false}
                    query={queryState.query}
                    filters={queryState.filters}
                    selectedFilter={queryState.selectedFilter}
                    pinnedTerms={queryState.pinnedTerms}
                    onChange={(query, pinnedTerms) => {
                      setQueryState({ ...queryState, query, pinnedTerms });
                    }}
                    onPin={pinnedTerms => {
                      setQueryState({ ...queryState, pinnedTerms });
                    }}
                    onFilterSelect={filterChangeHandler(
                      queryState,
                      setQueryState,
                    )}
                  />
                </Flex>
                <ExplorerTable
                  rows={data?.objects || []}
                  onColumnHeaderClick={columnHeaderHandler(
                    queryState,
                    setQueryState,
                  )}
                />
                <Flex wide center>
                  <Box p={2}>
                    <IconButton
                      disabled={queryState.offset === 0}
                      onClick={handlePageBack}
                    >
                      <Icon size={24} type={IconType.NavigateBeforeIcon} />
                    </IconButton>
                    <IconButton
                      disabled={
                        data?.objects &&
                        data?.objects?.length < queryState.limit
                      }
                      onClick={handlePageForward}
                    >
                      <Icon size={24} type={IconType.NavigateNextIcon} />
                    </IconButton>
                  </Box>
                </Flex>
              </>
            </RouterTab>
            <RouterTab name="Access Rules" path={`${Routes.Explorer}/access`}>
              <AccessRulesDebugger />
            </RouterTab>
          </SubRouterTabs>
        </div>
      </ContentWrapper>
    </PageTemplate>
  );
}

export default styled(Explorer).attrs({ className: Explorer.name })`
  td:last-child {
    white-space: pre-wrap;
    overflow-wrap: break-word;
    word-wrap: break-word;
  }
`;
